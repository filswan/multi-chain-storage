<template>
<div>
    <el-dialog :title="$t('uploadFile.nft_title')+'NFT'" :close-on-click-modal="false" :width="widthDia" :visible.sync="mineVisible"
        :before-close="closeDia">
        <div v-loading="hashload" :element-loading-text="isload?isloadText:''">
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
                <el-button class="cancel"
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
                    attributes: [{ trait_type: 'Size', value: parseInt(this.mintRow.file_size) }],
                    external_url: ''
                },
                rules: {
                    name: [
                        { required: true, message: ' ', trigger: 'blur' }
                    ],
                },
                widthDia: document.body.clientWidth<600?'90%':'500px',
                hashload: false,
                isload: false,
                isloadText: this.$t('uploadFile.payment_tip_deal'),
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
                            formData.append('duration', 525)
                            formData.append('file_type', 1)
                            formData.append('wallet_address', that.metaAddress)

                            const metadataUploadResponse = await that.sendPostRequest(`${process.env.BASE_PAYMENT_GATEWAY_API}api/v1/storage/ipfs/upload`, formData)
                            const nftUrl = metadataUploadResponse.data.ipfs_url
                            that.isloadText = that.$t('uploadFile.payment_tip_deal01')
                            console.log('upload success')
                            let nftContract = new web3.eth.Contract(
                                nftContractAbi,
                                that.$root.MINT_CONTRACT,
                                { from: that.metaAddress, gas: web3.utils.toHex(that.$root.PAY_GAS_LIMIT) },
                            )

                            const mintResult = await that.mintContract(nftContract, nftUrl)

                            // that.tokenId = await nftContract.methods.totalSupply().call()                            
                            console.log('totalSupply success', that.tokenId)
                            let mintInfoJson = {
                                source_file_upload_id: that.mintRow.source_file_upload_id,
                                payload_cid: that.mintRow.payload_cid,
                                tx_hash: that.nftHash,
                                token_id: that.tokenId,
                                mint_address: that.$root.MINT_CONTRACT
                            }
                            const mintInfoResponse = await that.sendPostRequest(`${process.env.BASE_PAYMENT_GATEWAY_API}api/v1/storage/mint/info`, mintInfoJson)
                            
                            if(mintInfoResponse) that.$emit('getMintDialog', false, that.tokenId, that.nftHash)
                        }
                    } else {
                        console.log('error submit!!');
                        return false;
                    }
                });
            },
            async mintContract(nftContract, nftUrl) {
                try {
                    console.log('start mint contract')
                    const transaction = await nftContract.methods
                    .mintData(that.metaAddress, nftUrl)
                    .send()
                    .on('transactionHash', function(hash){
                        that.nftHash = hash
                        that.isloadText = that.$t('uploadFile.payment_tip_deal02')
                        console.log('transactionHash console:', that.nftHash);
                    })
                    .on('receipt', function(receipt){
                        // receipt example
                        console.log('receipt console:', receipt);
                        return true
                    })
                    .on('error', function(error){
                        console.log('error console:', error)
                        that.hashload = false
                        that.isload = false
                        return false
                    });
                    // console.log('transaction.events:', transaction.events)
                    that.tokenId = transaction.events.Mint.returnValues.tokenId_
                    console.log('mintData success')
                } catch (err) {
                    console.log('err.response', err);
                    if (err.includes('not mined within 50 blocks')) {
                        const handle = setInterval(() => {
                            web3.eth.getTransactionReceipt(err.response.transactionHash).then((resp) => {
                                console.log('checking ... ');
                                if (resp != null && resp.blockNumber > 0) {
                                    clearInterval(handle);
                                    return true
                                }
                            });
                        },2000);
                    } else {
                        console.log(err.response);
                        return false
                    }
                }
            },
            closeDia() {
                that.$emit('getMintDialog', false)
            },
            async init(){
                that.hashload = true
                
                that.ruleForm.name = that.mintRow.file_name
                that.ruleForm.image = that.mintRow.ipfs_url
                that.ruleForm.external_url = that.mintRow.ipfs_url

                const hashRes = await that.sendRequest(`${process.env.BASE_PAYMENT_GATEWAY_API}api/v1/billing/deal/lockpayment/info?payload_cid=${that.mintRow.payload_cid}&wallet_address=${that.metaAddress}&source_file_upload_id=${that.mintRow.source_file_upload_id}`)
                that.ruleForm.tx_hash = hashRes.data.pay_tx_hash

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
    display: flex;
    align-items: center;
    justify-content: center;
    .el-dialog{
        background: #fff;
        margin: auto !important;
        box-shadow: 0 0 13px rgba(128,128,128,0.8);
        border-radius: 0.2rem;
        .el-dialog__header{
            padding: 0.3rem 0.4rem;
            display: flex;
            border-bottom: 1px solid #dfdfdf;
            .el-dialog__title{
                color: #000;
                font-size: 0.22rem;
                font-weight: 500;
                line-height: 1;
                text-transform: capitalize;
            }
            .el-dialog__headerbtn{
                display: none;
            }
        }
        .el-dialog__body{
            padding: 0.2rem 0.4rem;
            .el-form{
                width: 100%; 
                margin: auto; 
                justify-content: flex-start;
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
                        font-size: 0.2rem;
                    }
                    .el-form-item__content{
                        width: 100%;
                        display: flex;
                        flex-wrap: wrap;
                        overflow: hidden;
                        font-size: 0.2rem;
                        color: #555;
                        .el-input, .el-textarea{
                            margin: 0 5px 0 0;
                            .el-input__inner, .el-textarea__inner{
                                width: 100%;
                                font-size: 0.2rem;
                                font-family: inherit;
                                color: #555;
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
            justify-content: space-between;
            align-items: center;
            width: 100%;
            margin: 0.25rem auto 0;
            .el-button{
                height: 0.6rem;
                padding: 0;
                margin-left: 0;
                line-height: 0.6rem;
                font-size: 0.22rem;
                font-family: inherit;
                color: #fff;
                border: 0;
                background: linear-gradient(45deg,#4f8aff, #4b5eff);
                border-radius: 14px;
                width: calc(50% - 0.15rem);
                &:hover{
                    opacity: .9;
                }
            }
            .cancel{
                background: #dadada;
                transition: background-color .3s;
                &:hover{
                    background: linear-gradient(45deg,#4f8aff, #4b5eff);
                }
            }
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
