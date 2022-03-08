<template>
<div>
    <el-dialog :title="$t('uploadFile.nft_title')+'NFT'" :width="widthDia" :visible.sync="mineVisible"
        :before-close="closeDia">
        <div v-loading="hashload" :element-loading-text="isload?$t('uploadFile.payment_tip_deal'):''">
            <el-form :model="ruleForm" :rules="rules" ref="ruleForm">
                <el-form-item :label="'NFT '+$t('uploadFile.nft_Name')" prop="name">
                    <el-input v-model="ruleForm.name" placeholder=""></el-input>
                </el-form-item>
                <el-form-item :label="'NFT '+$t('uploadFile.nft_Description')" prop="description">
                    <el-input v-model="ruleForm.description" type="textarea" :rows="2"></el-input>
                </el-form-item>
                <el-form-item :label="$t('uploadFile.nft_IPFSURL')" prop="image">
                    <el-input v-model="ruleForm.image" readOnly></el-input>
                </el-form-item>
                <el-form-item :label="$t('uploadFile.Payment_Transaction_Hash')" prop="tx_hash">
                    <el-input v-model="ruleForm.tx_hash" readOnly></el-input>
                </el-form-item>
                <el-form-item :label="$t('uploadFile.file_size')" prop="attributes">
                    <el-input v-model="ruleForm.attributes[0].value" readOnly></el-input>
                </el-form-item>
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button 
                    type="info" @click="closeDia">{{$t('uploadFile.Back')}}</el-button>
                <el-button 
                    type="primary" @click="submitForm('ruleForm')">{{isload ? $t('uploadFile.Minting') : $t('uploadFile.Mint_NFT')}}</el-button>
            </div>
        </div>
    </el-dialog>
</div>
</template>

<script>
    let that
    import axios from 'axios'
    import nftContractAbi from '@/utils/DatabaseMinter.json'
    export default {
        name: "mint_tip",
        data() {
            return {
                ruleForm: {
                    name: '',
                    image: '',
                    description: '',
                    tx_hash: '',
                    attributes: [{ trait_type: 'Size', value: parseInt(this.mintRow.file_size) }]
                },
                rules: {
                    name: [
                        { required: true, message: ' ', trigger: 'blur' }
                    ],
                },
                widthDia: document.body.clientWidth<600?'90%':'500px',
                hashload: false,
                isload: false,
                nftHash: '',
                tokenId: ''
            };
        },
        props: ['mintRow', 'mineVisible'],
        components: {},
        computed: {
            metaAddress() {
                return this.$store.getters.metaAddress
            }
        },
        methods: {
            submitForm(formName) {
                that.$refs[formName].validate(async (valid) => {
                    if (valid) {
                        if(that.metaAddress){
                            that.hashload = true
                            that.isload = true

                            const fileBlob = new Blob([JSON.stringify(that.ruleForm)], {
                                type: 'application/json',
                            })

                            var formData = new FormData()
                            formData.append('file', fileBlob, `${that.ruleForm.name}.json`)
                            formData.append('duration', 180)
                            formData.append('file_type', 1)
                            formData.append('wallet_address', that.metaAddress)

                            const metadataUploadResponse = await that.sendPostRequest(`${process.env.BASE_PAYMENT_GATEWAY_API}api/v1/storage/ipfs/upload`, formData)
                            const nftUrl = metadataUploadResponse.data.ipfs_url

                            let nftContract = new web3.eth.Contract(
                                nftContractAbi,
                                that.$root.MINT_CONTRACT,
                                { from: that.metaAddress, gas: web3.utils.toHex(that.$root.PAY_GAS_LIMIT) },
                            )
                            const transaction = await nftContract.methods
                            .mintData(that.metaAddress, nftUrl)
                            .send()
                            .on('transactionHash', function(hash){
                                that.nftHash = hash
                                console.log('transactionHash console:', that.nftHash);
                            })
                            .on('receipt', function(receipt){
                                // receipt example
                                // console.log('receipt console:', receipt);
                            })
                            .on('error', function(error){
                                console.log('error console:', error)
                                that.hashload = false
                                that.isload = false
                            });

                            that.tokenId = await nftContract.methods.totalSupply().call()                            
                            
                            let mintInfoJson = {
                                payload_cid: that.mintRow.payload_cid,
                                tx_hash: that.nftHash,
                                token_id: that.tokenId,
                                mint_address: that.$root.MINT_CONTRACT
                            }
                            const mintInfoResponse = await that.sendPostRequest(`${process.env.BASE_PAYMENT_GATEWAY_API}api/v1/storage/mint/info`, mintInfoJson)
                            
                            that.$emit('getMintDialog', false, that.tokenId, that.nftHash)
                        }
                    } else {
                        console.log('error submit!!');
                        return false;
                    }
                });
            },
            closeDia() {
                that.$emit('getMintDialog', false)
            },
            async init(){
                that.hashload = true
                
                that.ruleForm.name = that.mintRow.file_name
                that.ruleForm.image = that.mintRow.ipfs_url

                const hashRes = await that.sendRequest(`${process.env.BASE_PAYMENT_GATEWAY_API}api/v1/billing/deal/lockpayment/info?payload_cid=${that.mintRow.payload_cid}&wallet_address=${that.metaAddress}`)
                that.ruleForm.tx_hash = hashRes.data.tx_hash

                that.hashload = false
            },
            async sendRequest(apilink) {
                try {
                    const response = await axios.get(apilink)
                    return response.data
                } catch (err) {
                    console.error(err)
                }
            },
            async sendPostRequest(apilink, jsonObject) {
                try {
                    const response = await axios.post(apilink, jsonObject)
                    return response.data
                } catch (err) {
                    console.error(err)
                }
            },
        },
        mounted() {
            that = this
            that.init()
        },
        watch: {},
        filters: {}
    };
</script>


<style scoped lang="scss">
.el-dialog__wrapper /deep/{
    .el-dialog{
        box-shadow: 0 0 13px rgba(128,128,128,0.8);
        .el-dialog__header{
            padding: 0.2rem 0.2rem 0.15rem;
            display: flex;
            border-bottom: 1px solid #eee;
            .el-dialog__title{
                color: #222222;
                font-size: 0.18rem;
                font-weight: 600;
            }
        }
        .el-dialog__body{
            padding: 0.1rem 0.2rem 0.2rem;
            .el-form{
                .err{
                    .el-form-item__label{
                    color: red;
                    }
                    .el-input{
                    .el-input__inner{
                        border-color: red;
                    }
                    }
                }
                .el-form-item{
                    display: flex;
                    align-items: center;
                    flex-wrap: wrap;
                    margin-bottom: 0.05rem;
                    .el-form-item__label{
                        width: 100%;
                        color: #000000;
                        line-height: 2.5;
                        word-break: break-word;
                        text-align: left;
                    }
                    .el-form-item__content{
                        width: 100%;
                        display: flex;
                        flex-wrap: wrap;
                        overflow: hidden;
                        font-size: 0.14rem;
                        .el-input{
                            margin: 0 5px 0 0;
                            .el-input__inner{
                                width: 100%;
                            }
                            .el-input__inner[readOnly]{
                                background: #f9f9f9;
                            }
                        }
                    }
                    p{
                        width: 100%;
                        margin: 0.05rem 0;
                        font-size: 0.16rem;
                        font-weight: 100;
                        color: red;
                        white-space: normal;
                        word-break: break-all;
                        line-height: 1;
                    }
                }
            }
        }
        .dialog-footer{
            display: flex;
            align-items: center;
            justify-content: flex-end;
            margin: 0.15rem 0 0;
        }
    }
}
@media screen and (max-width: 599px){
    .el-dialog__wrapper /deep/{
        .el-dialog{
            .el-dialog__header{
                .el-dialog__title{
                    font-size: 0.16rem;
                }
            }
            .el-dialog__body{
                padding: 0.15rem;
                .el-form{
                    .el-form-item{
                        display: flex;
                        flex-wrap: wrap;
                        margin: 0 0 0.05rem;
                        .el-form-item__label{
                            width: 100%;
                            margin: 0;
                            text-align: left;
                        }
                        .el-form-item__content{
                            width: 100%;
                            margin: 0;
                        }
                    }
                }
            }
        }
    }
}
</style>
