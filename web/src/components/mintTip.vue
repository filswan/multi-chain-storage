<template>
  <div>
    <el-dialog :close-on-click-modal="false" :width="widthDia" :visible.sync="mineVisible" :before-close="closeDia">
      <div slot="title" class="dialog-header">
        <span class="el-dialog__title">{{$t('uploadFile.nft_title')+'NFT'}}</span>
        <el-button type="primary" size="mini" v-if="mintIndex === 'list' && !hashload" @click="init">{{$t('my_profile.swan_miners_Refresh')}}</el-button>
      </div>
      <div v-if="mintIndex === 'list'" v-loading="hashload">
        <div class="mint_body">
          <el-card shadow="always" class="mint_card" v-for="(nftMint, n) in nftMintData" :key="n">
            <div class="mint_flex">
              <div class="details">
                <div class="img">
                  <img :src="nftMint.image_url" v-if="nftMint.image_url" />
                </div>
                <span>{{nftMint.name || nftMint.address||'-'}}</span>
              </div>
              <el-button v-if="nftMint.address && nftMint.address.indexOf('0x') > -1" type="primary" size="mini" @click="handleMint(nftMint, 'mint')">Mint here</el-button>
              <el-popover v-else placement="top" popper-class="elPopTitle" class="item" width="200" trigger="hover" :content="$t('uploadFile.mint_pending')">
                <el-button slot="reference" disabled type="primary" size="mini" @click="handleMint(nftMint, 'mint')">Mint here</el-button>
              </el-popover>
            </div>
          </el-card>
          <el-card shadow="always" class="mint_card" v-for="(nftView, v) in nftViewData" :key="v+nftMintData.length">
            <div class="mint_flex">
              <div class="details">
                <div class="img">
                  <img :src="nftView.nft_collection_image_url" v-if="nftView.nft_collection_image_url" />
                </div>
                <span>{{nftView.nft_collection_name||nftView.nft_collection_address || nftView.mint_address || '-'}}</span>
              </div>
              <el-button type="primary" size="mini" @click="handleMint(nftView, 'view')">
                {{ $t('uploadFile.mint_view') }}
              </el-button>
            </div>
          </el-card>
          <p class="mint_nodata" v-if="nftMintData.length<1 && nftViewData.length<1">{{$t('uploadFile.nft_list_empty')}}</p>
        </div>

        <div slot="footer" class="dialog-footer">
          <el-button class="cancel" type="info" @click="closeDia">{{ $t('uploadFile.Back') }}</el-button>
          <el-button type="primary" @click="mintIndex = 'create'">
            {{ isload ? $t('uploadFile.Minting') : $t('uploadFile.Mint_Create_NFT') }}
          </el-button>
        </div>
      </div>
      <div v-else-if="mintIndex === 'mint'" v-loading="hashload" :element-loading-text="isload?isloadText:''">
        <el-form :model="ruleForm" :rules="rules" ref="ruleForm">
          <el-form-item :label="'NFT '+$t('uploadFile.nft_Name')" prop="name" class="flex_float">
            <el-input v-model="ruleForm.name" placeholder=""></el-input>
          </el-form-item>
          <el-form-item :label="'NFT '+$t('uploadFile.nft_Amount')" prop="amount" class="flex_float">
            <el-input-number v-model="ruleForm.amount" controls-position="right" :min="1" placeholder=""></el-input-number>
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
          <el-button class="cancel" type="info" @click="mintIndex = 'list'">{{ $t('uploadFile.Back') }}</el-button>
          <el-button type="primary" @click="submitForm('ruleForm', 'mint')">{{ $t('uploadFile.Mint_NFT') }}</el-button>
        </div>
      </div>
      <div v-else-if="mintIndex === 'create'" v-loading="hashload" :element-loading-text="isload?isloadText:''">
        <el-form :model="ruleCreateForm" :rules="rules" ref="ruleCreateForm">
          <el-form-item :label="'NFT '+$t('uploadFile.nft_Image')" prop="images">
            <el-upload class="upload-demo" ref="uploadFileRef" action="customize" drag list-type="picture-card" :auto-upload="false" :file-list="ruleCreateForm.images" :on-change="handleChange" :on-remove="handleRemove">
              <i class="el-icon-plus"></i>
            </el-upload>
          </el-form-item>
          <el-form-item :label="'NFT '+$t('uploadFile.nft_Name')" prop="name">
            <el-input v-model="ruleCreateForm.name" placeholder=""></el-input>
          </el-form-item>
          <el-form-item :label="'NFT '+$t('uploadFile.nft_Description')" prop="description">
            <el-input v-model="ruleCreateForm.description" type="textarea" :rows="2"></el-input>
          </el-form-item>
          <el-form-item :label="$t('uploadFile.nft_External')" prop="external_link">
            <el-input v-model="ruleCreateForm.external_link"></el-input>
          </el-form-item>
          <el-form-item :label="$t('uploadFile.nft_Seller')" prop="seller_fee_basis_points">
            <el-input-number v-model="ruleCreateForm.seller_fee_basis_points" controls-position="right" :min="0" :max="10000" placeholder=""></el-input-number>
          </el-form-item>
          <el-form-item :label="$t('uploadFile.nft_Fee_Recipient')" prop="fee_recipient">
            <el-input v-model="ruleCreateForm.fee_recipient"></el-input>
          </el-form-item>
        </el-form>

        <div slot="footer" class="dialog-footer">
          <el-button class="cancel" type="info" @click="mintIndex = 'list'">{{ $t('uploadFile.Back') }}</el-button>
          <el-button type="primary" @click="submitForm('ruleCreateForm', 'create')">
            {{ isload ? $t('uploadFile.Minting') : $t('uploadFile.Mint_Create_NFT') }}
          </el-button>
        </div>
      </div>
    </el-dialog>
  </div>
</template>

<script>
// import axios from 'axios'
import CollectionFactoryAbi from '@/utils/CollectionFactory.json'

let that
export default {
  name: 'mint_tip',
  data () {
    return {
      ruleForm: {
        name: '',
        image: '',
        description: '',
        tx_hash: '',
        attributes: [{ trait_type: 'Size', value: parseInt(this.mintRow.file_size) }],
        external_url: '',
        amount: 1,
        nftMint: null
      },
      ruleCreateForm: {
        name: '',
        description: '',
        tx_hash: '',
        attributes: [{ trait_type: 'Size', value: parseInt(this.mintRow.file_size) }],
        images: [],
        external_link: '',
        seller_fee_basis_points: 0,
        fee_recipient: this.metaAddress,
        file: '',
        fileRaw: {},
        collection_name: ''
      },
      rules: {
        name: [
          { required: true, message: ' ', trigger: 'blur' }
        ]
      },
      widthDia: document.body.clientWidth > 1600 ? '550px' : document.body.clientWidth > 600 ? '450px' : '90%',
      hashload: false,
      isload: false,
      isloadText: this.$t('uploadFile.payment_tip_deal'),
      nftHash: '',
      tokenId: '',
      mintIndex: 'list',
      mintCollectionAddress: {},
      nftMintData: [],
      nftViewData: []
    }
  },
  props: ['mintRow', 'mineVisible'],
  components: {},
  computed: {
    metaAddress () {
      return this.$store.getters.metaAddress
    }
  },
  methods: {
    handleChange (file, fileList) {
      if (fileList.length > 0) that.ruleCreateForm.images = [fileList[fileList.length - 1]]
      that.uploadFileImage(file, fileList)
    },
    handleRemove (file, fileList) {
      // console.log(file, fileList);
      that.ruleCreateForm.images = []
    },
    async uploadFileImage (file, fileList) {
      // const isJPG = file.raw.type === 'image/jpeg' || file.raw.type === 'image/png'
      const isJPG = file.raw.type.indexOf('image') > -1
      const isLt2M = file.size / 1024 / 1024 / 1024 <= 25 // or 1000
      if (!isJPG) {
        that.$message.error('Uploaded avatar images can only be in JPG or png format!')
        that.ruleCreateForm.images = []
        return false
      }
      if (!isLt2M) {
        that.$message.error(that.$t('deal.upload_form_file_tip'))
        that.ruleCreateForm.images = []
        return false
      }
      let reg = new RegExp(' ', 'g')
      if (file.name.indexOf(' ') > -1) file.name = file.name.replace(reg, '_')
      that.ruleCreateForm.fileRaw = file
    },
    async uploadImage (raw, name) {
      var imageData = new FormData()
      imageData.append('file', raw, name)
      imageData.append('duration', '525')
      imageData.append('file_type', '1')
      imageData.append('wallet_address', that.metaAddress)
      const imageResponse = await that.sendPostRequest(`${that.baseAPIURL}api/v1/storage/ipfs/upload`, imageData)
      if (imageResponse && imageResponse.status === 'success') return imageResponse.data.ipfs_url || ''
      else {
        that.$message.error(imageResponse.message || 'Failed!')
        that.closeDia()
        return ''
      }
    },
    handleMint (row, type) {
      that.mintCollectionAddress = row
      if (type === 'view') that.$emit('getMintDialog', false, row)
      else that.mintIndex = 'mint'
    },
    submitForm (formName, type) {
      that.$refs[formName].validate(async (valid) => {
        if (valid) {
          if (that.metaAddress) {
            that.hashload = true
            that.isload = true

            const fileBlob = new Blob([JSON.stringify(type === 'create' ? that.ruleCreateForm : that.ruleForm)], {
              type: 'application/json'
            })
            const nftUrl = await that.uploadImage(fileBlob, `${type === 'create' ? that.ruleCreateForm.name : that.ruleForm.name}.json`)
            if (!nftUrl) return
            console.log('upload success')

            let CollectionFactory = new that.$web3Init.eth.Contract(
              CollectionFactoryAbi,
              that.$root.COLLECTION_FACTORY_ADDRESS,
              // '0xeC8988d45F1E85Fb9cF816f365Bf5775a78dcA3a',
              { from: that.metaAddress, gas: that.$web3Init.utils.toHex(that.$root.PAY_GAS_LIMIT) }
            )

            if (type === 'create') {
              CollectionFactory.methods.createCollection(that.ruleCreateForm.name, nftUrl)
                .send()
                .on('transactionHash', async (hash) => {
                  console.log('hash console:', hash)
                  if (that.ruleCreateForm.fileRaw.raw) {
                    that.ruleCreateForm.file = await that.uploadImage(that.ruleCreateForm.fileRaw.raw, that.ruleCreateForm.fileRaw.name)
                  }
                  let mintInfoJson = {
                    tx_hash: hash,
                    name: that.ruleCreateForm.name,
                    description: that.ruleCreateForm.description,
                    image_url: that.ruleCreateForm.file,
                    external_link: that.ruleCreateForm.external_link,
                    seller_fee: that.ruleCreateForm.seller_fee_basis_points || 0,
                    wallet_recipient: that.ruleCreateForm.fee_recipient
                  }
                  await that.sendPostRequest(`${that.baseAPIURL}api/v1/storage/mint/nft_collection`, mintInfoJson)

                  that.mintCollectionAddress = {}
                  that.mintIndex = 'list'
                  that.isload = false
                  that.init()
                  that.$refs['ruleCreateForm'].resetFields()
                })
                .on('confirmation', function (confirmationNumber, receipt) {
                  // console.log('confirmationNumber console:', confirmationNumber, receipt)
                })
                .on('receipt', function (receipt) {
                  // receipt example
                  console.log('create receipt console:', receipt)
                })
            } else {
              that.isloadText = that.$t('uploadFile.payment_tip_deal01')
              await that.mintContract(CollectionFactory, that.mintCollectionAddress.address, nftUrl)
              if (that.tokenId) {
                console.log('totalSupply success', that.tokenId)
                let mintInfoJson = {
                  source_file_upload_id: that.mintRow.source_file_upload_id,
                  payload_cid: that.mintRow.payload_cid,
                  tx_hash: that.nftHash,
                  nft_collection_id: that.mintCollectionAddress.id,
                  token_id: parseInt(that.tokenId),
                  mint_address: that.mintCollectionAddress.address,
                  name: that.ruleForm.name,
                  description: that.ruleForm.description
                }
                const mintInfoResponse = await that.sendPostRequest(`${that.baseAPIURL}api/v1/storage/mint/info`, mintInfoJson)
                if (mintInfoResponse) that.$emit('getMintDialog', false, mintInfoJson)
              }
              that.hashload = false
              that.isload = false
            }
          }
        } else {
          console.log('error submit!!')
          return false
        }
      })
    },
    async mintContract (CollectionFactory, collectAddress, nftUrl) {
      try {
        console.log('start mint contract', collectAddress)
        let payObject = {
          gasPrice: await that.$web3Init.eth.getGasPrice()
        }
        const transaction = await CollectionFactory.methods
          .mint(collectAddress, that.metaAddress, that.ruleForm.amount || 1, nftUrl)
          .send(payObject)
          .on('transactionHash', async function (hash) {
            that.nftHash = hash
            that.isloadText = that.$t('uploadFile.payment_tip_deal02')
            console.log('transactionHash console:', that.nftHash)
          })
          .on('receipt', function (receipt) {
            // receipt example
            console.log('receipt console:', receipt)
            return true
          })
          .on('error', function (error) {
            console.log('error console:', error)
            that.hashload = false
            that.isload = false
            return false
          })
        // console.log('transaction.events:', transaction)
        that.tokenId = transaction.events.TransferSingle.returnValues.id
        console.log('mintUnique success')
      } catch (err) {
        console.log('err.response', err.message)
        if (err.message.includes('not mined within 50 blocks')) {
          const handle = setInterval(() => {
            that.$web3Init.eth.getTransactionReceipt(err.response.transactionHash).then((resp) => {
              console.log('checking ... ')
              if (resp != null && resp.blockNumber > 0) {
                clearInterval(handle)
                return true
              }
            })
          }, 2000)
        } else {
          console.log(err.response)
          return false
        }
      }
    },
    closeDia () {
      that.$emit('getMintDialog', false)
    },
    async init () {
      that.hashload = true
      that.nftMintData = []
      that.nftViewData = []
      that.ruleForm.name = that.mintRow.file_name
      that.ruleForm.image = that.mintRow.ipfs_url
      that.ruleForm.external_url = that.mintRow.ipfs_url

      const hashRes = await that.sendRequest(`${that.baseAPIURL}api/v1/billing/deal/lockpayment/info?payload_cid=${that.mintRow.payload_cid}&wallet_address=${that.metaAddress}&source_file_upload_id=${that.mintRow.source_file_upload_id}`)
      that.ruleForm.tx_hash = hashRes.data.pay_tx_hash

      const nftMintRes = await that.sendRequest(`${that.baseAPIURL}api/v1/storage/mint/nft_collections`)
      if (nftMintRes && nftMintRes.status === 'success') that.nftMintData = nftMintRes.data || []
      else {
        that.closeDia()
        return
      }

      const nftViewRes = await that.sendRequest(`${that.baseAPIURL}api/v1/storage/mint/info/${that.mintRow.source_file_upload_id}`)
      if (nftViewRes && nftViewRes.status === 'success') that.nftViewData = nftViewRes.data || []
      else {
        that.closeDia()
        return
      }

      that.nftViewData.forEach(element => {
        that.nftMintData.forEach((mint, m) => {
          if (mint.address === element.nft_collection_address) that.nftMintData.splice(m, 1)
        })
      })
      that.nftMintData.sort((a, b) => b.is_default - a.is_default)

      that.hashload = false
    },
    async sendRequest (apilink) {
      try {
        const response = await axios.get(apilink, {
          headers: {
            'Authorization': 'Bearer ' + that.$store.getters.mcsjwtToken
          }
        })
        return response.data
      } catch (err) {
        console.error(err, err.response)
        if (err.response) {
          that.$message.error(err.response.data.message || 'Failed!')
          return err.response.data
        }
      }
    },
    async sendPostRequest (apilink, jsonObject) {
      try {
        const response = await axios.post(apilink, jsonObject, {
          headers: {
            'Authorization': 'Bearer ' + that.$store.getters.mcsjwtToken
          }
        })
        return response.data
      } catch (err) {
        console.error(err)
      }
    }
  },
  mounted () {
    that = this
    that.ruleCreateForm.fee_recipient = that.metaAddress
    that.init()
  },
  watch: {},
  filters: {}
}
</script>

<style scoped lang="scss">
.el-dialog__wrapper /deep/ {
  display: flex;
  align-items: center;
  justify-content: center;

  .el-dialog {
    background: #fff;
    margin: auto !important;
    box-shadow: 0 0 13px rgba(128, 128, 128, 0.8);
    border-radius: 0.2rem;

    .el-dialog__header {
      padding: 0.25rem 0.4rem;
      display: flex;
      border-bottom: 1px solid #dfdfdf;
      .dialog-header {
        display: flex;
        align-items: center;
        justify-content: space-between;
        width: 100%;
        .el-button {
          padding: 7px 10px;
          font-size: 14px;
          font-family: inherit;
          color: #fff;
          border-radius: 0.1rem;
          @media screen and (max-width: 1600px) {
            font-size: 13px;
          }
          @media screen and (max-width: 768px) {
            font-size: 12px;
          }

          &:hover {
            opacity: 0.9;
          }
        }
      }
      .el-dialog__title {
        color: #000;
        font-size: 0.22rem;
        font-weight: 500;
        line-height: 1;
        text-transform: capitalize;
      }

      .el-dialog__headerbtn {
        display: none;
      }
    }

    .el-dialog__body {
      min-height: 250px;
      padding: 0.2rem 0.4rem;

      .el-form {
        width: 100%;
        margin: auto;
        justify-content: flex-start;

        .err {
          .el-form-item__label {
            color: red;
          }

          .el-input {
            .el-input__inner {
              border-color: red;
            }
          }
        }

        .el-form-item {
          display: flex;
          align-items: center;
          flex-wrap: wrap;
          width: 100%;
          margin-bottom: 0.05rem;

          .el-form-item__label {
            width: 100%;
            color: #000000;
            line-height: 2;
            word-break: break-word;
            text-align: left;
            font-size: 0.2rem;

            .important {
              color: #f56c6c;
              margin-right: 4px;
            }
          }

          .el-form-item__content {
            width: 100%;
            display: flex;
            flex-wrap: wrap;
            overflow: hidden;
            font-size: 0.2rem;
            color: #555;

            .el-input,
            .el-textarea {
              margin: 0 5px 0 0;

              .el-input__inner,
              .el-textarea__inner {
                width: 100%;
                font-size: 0.2rem;
                font-family: inherit;
                color: #555;
                text-align: left;
              }

              .el-input__inner[readOnly] {
                background: #f9f9f9;
              }
            }

            .el-select,
            .el-input-number {
              width: 100%;

              .el-input {
                .el-input__inner[readOnly] {
                  background: transparent;
                }
              }
            }

            .upload-demo {
              .el-upload--text {
                width: auto;
                height: auto;
                border: 0;
              }
            }

            .el-upload--picture-card,
            .el-upload-list--picture-card .el-upload-list__item {
              width: 65px;
              height: 65px;
              line-height: 70px;
              margin-bottom: 0;

              .el-upload-dragger {
                width: 100%;
                height: 100%;
              }
            }
          }

          p {
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

        .flex_float {
          float: left;
          width: 50%;

          .el-form-item__content {
            display: flex;
          }
        }
      }

      .mint_body {
        width: 100%;
        padding: 0 0.4rem;
        margin-left: -0.4rem;
        margin-bottom: 0.5rem;
        max-height: 250px;
        min-height: 180px;
        overflow: hidden;
        overflow-y: scroll;

        &::-webkit-scrollbar-track {
          background: transparent;
        }

        &::-webkit-scrollbar {
          width: 6px;
          background: transparent;
        }

        &::-webkit-scrollbar-thumb {
          background: #ccc;
        }

        .mint_nodata {
          text-align: center;
        }

        .mint_card {
          width: 100%;
          margin: 0.2rem auto;

          .el-card__body {
            padding: 0;

            .mint_flex {
              display: flex;
              align-items: center;
              justify-content: space-between;
              padding: 0.1rem;

              .details {
                display: flex;
                align-items: center;
                flex-wrap: wrap;
                width: calc(100% - 100px);
                padding: 0;
                line-height: 1.2;
                .img {
                  display: block;
                  width: 40px;
                  height: 40px;
                  margin: 0 7px 0 0;
                  border: 1px solid #4b5eff;
                  border-radius: 100%;
                  img {
                    display: block;
                    width: 100%;
                    height: 100%;
                    border-radius: 100%;
                  }
                }

                span {
                  width: calc(100% - 50px);
                  overflow: hidden;
                  text-overflow: ellipsis;
                  white-space: normal;
                  display: -webkit-box;
                  -webkit-line-clamp: 1;
                  -webkit-box-orient: vertical;
                }
              }

              .el-button {
                padding: 7px 10px;
                font-size: 14px;
                font-family: inherit;
                color: #fff;
                border-radius: 0.1rem;
                @media screen and (max-width: 1600px) {
                  font-size: 13px;
                }
                @media screen and (max-width: 768px) {
                  font-size: 12px;
                }

                &:hover {
                  opacity: 0.9;
                }
              }
            }
          }
        }

        .view_style {
          cursor: pointer;

          .el-card__body {
            .mint_flex {
              .details {
                width: 100%;
              }
            }
          }
        }
      }
    }

    .dialog-footer {
      display: flex;
      justify-content: space-between;
      align-items: center;
      width: 100%;
      margin: 0.25rem auto 0;

      .el-button {
        height: 0.6rem;
        padding: 0;
        margin-left: 0;
        line-height: 0.6rem;
        font-size: 0.22rem;
        font-family: inherit;
        color: #fff;
        border: 0;
        background: linear-gradient(45deg, #4f8aff, #4b5eff);
        border-radius: 14px;
        width: calc(65% - 0.15rem);

        &:hover {
          opacity: 0.9;
        }
      }

      .cancel {
        width: calc(35% - 0.15rem);
        background: #dadada;
        transition: background-color 0.3s;

        &:hover {
          background: linear-gradient(45deg, #4f8aff, #4b5eff);
        }
      }
    }
  }
}

@media screen and (max-width: 599px) {
  .el-dialog__wrapper /deep/ {
    .el-dialog {
      .el-dialog__header {
        .el-dialog__title {
          font-size: 0.16rem;
        }
      }

      .el-dialog__body {
        padding: 0.15rem;

        .el-form {
          .el-form-item {
            display: flex;
            flex-wrap: wrap;
            margin: 0 0 0.05rem;

            .el-form-item__label {
              width: 100%;
              margin: 0;
              text-align: left;
            }

            .el-form-item__content {
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
