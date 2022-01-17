<template>
    <div id="Create">
        <div class="upload" v-loading="loading">
            <div class="upload_title">{{$t('uploadFile.uploadFile_title')}}</div>
            <div class="upload_form">
                <el-form :model="ruleForm" status-icon :rules="rules" ref="ruleForm" class="demo-ruleForm">
                    <el-form-item prop="fileList" :label="$t('uploadFile.upload')">
                        <div>
                            <el-upload
                                class="upload-demo"
                                ref="uploadFileRef"
                                action="customize"
                                :http-request="uploadFile"
                                :file-list="ruleForm.fileList"
                                :on-change="handleChange"
                                :on-remove="handleRemove">
                                <el-button size="small" type="primary" icon="el-icon-plus">{{$t('uploadFile.upload')}}</el-button>
                            </el-upload>
                            <p v-if="ruleForm.fileList.length>0" style="display: flex;align-items: center;">
                                <svg t="1637031488880" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="3310" style="width: 0.13rem;height: 0.13rem;margin: 0 7px 0 5px;"><path d="M512 1024a512 512 0 1 1 512-512 32 32 0 0 1-32 32h-448v448a32 32 0 0 1-32 32zM512 64a448 448 0 0 0-32 896V512a32 32 0 0 1 32-32h448A448 448 0 0 0 512 64z" fill="#999999" p-id="3311"></path><path d="M858.88 976a32 32 0 0 1-32-32V640a32 32 0 0 1 32-32 32 32 0 0 1 32 32v304a32 32 0 0 1-32 32z" fill="#999999" p-id="3312"></path><path d="M757.12 773.12a34.56 34.56 0 0 1-22.4-8.96 32 32 0 0 1 0-45.44l101.12-101.12a32 32 0 0 1 45.44 0 30.72 30.72 0 0 1 0 44.8l-101.12 101.76a34.56 34.56 0 0 1-23.04 8.96z" fill="#999999" p-id="3313"></path><path d="M960 773.12a32 32 0 0 1-22.4-8.96l-101.76-101.76a32 32 0 0 1 0-44.8 32 32 0 0 1 45.44 0l101.12 101.12a32 32 0 0 1-22.4 54.4z" fill="#999999" p-id="3314"></path></svg>
                                {{ruleForm.file_size}}
                            </p>
                        </div>
                    </el-form-item>
                    <el-form-item prop="duration">
                        <template slot="label">
                            {{$t('uploadFile.Duration')}} 
                            
                            <el-tooltip effect="dark" :content="$t('uploadFile.Duration_tooltip')" placement="top">
                                <img src="@/assets/images/info.png"/>
                            </el-tooltip>
                        </template>
                        <el-input-number v-model="ruleForm.duration" @change="calculation" controls-position="right" :min="180" :max="540" :step-strictly="true"></el-input-number> &nbsp; {{$t('components.day')}}<small> {{$t('components.interval')}}</small>
                        <!-- <el-input v-model="ruleForm.duration" type="number" style="max-width:130px"></el-input> &nbsp; {{$t('components.day')}} <small> {{$t('components.interval')}}</small> -->
                    </el-form-item>
                    <el-form-item prop="storage_cost">
                        <template slot="label">
                            {{$t('uploadFile.Estimated_Storage_Cost')}} 
                            
                            <el-tooltip effect="dark" :content="$t('uploadFile.Estimated_Storage_Cost_tooltip')" placement="top">
                                <img src="@/assets/images/info.png"/>
                            </el-tooltip>
                        </template>
                        <span  style="color:#ce2f21">{{ruleForm.storage_cost | NumStorage}} FIL</span> 
                    </el-form-item>
                </el-form>
                <div class="upload_plan">
                    <div class="title" :style="{'color': ruleForm.lock_plan_tip? '#f67e7e' : '#000'}">
                        {{$t('uploadFile.Select_Lock_Funds_Plan')}}
                        <el-tooltip effect="dark" :content="$t('uploadFile.Select_Lock_Funds_Plan_tooltip')" placement="top">
                            <img src="@/assets/images/info.png"/>
                        </el-tooltip>
                    </div>
                    <div class="desc">{{$t('uploadFile.latest_exchange_rate')}} {{biling_price}}.</div>
                    <div class="upload_plan_radio">
                        <el-radio-group v-model="ruleForm.lock_plan" @change="agreeChange">
                            <el-radio label="1" border>
                                <div class="title">{{$t('uploadFile.Low')}}</div>
                                <div class="cont">
                                    {{storage_cost_low}} <br/> USDC
                                </div>
                            </el-radio>
                            <el-radio label="2" border>
                                <div class="title">{{$t('uploadFile.Average')}}</div>
                                <div class="cont">
                                    {{storage_cost_average}} <br/> USDC
                                </div>
                            </el-radio>
                            <el-radio label="3" border>
                                <div class="title">{{$t('uploadFile.High')}}</div>
                                <div class="cont">
                                    {{storage_cost_high}} <br/> USDC
                                </div>
                            </el-radio>
                        </el-radio-group>
                    </div>
                </div>
                <div class="upload_bot">
                    <el-button type="primary" @click="submitForm('ruleForm')">{{$t('deal.Submit')}}</el-button>
                </div>
            </div>

            <div class="loadMetamaskPay" v-if="loadMetamaskPay">
                <div>
                    <div class="el-loading-spinner"><svg viewBox="25 25 50 50" class="circular"><circle cx="50" cy="50" r="20" fill="none" class="path"></circle></svg><!----></div>
                    <p>{{$t('uploadFile.payment_tip')}}</p>
                </div>
            </div>
        </div>

        <el-dialog title="" :visible.sync="finishTransaction" :width="width"
            custom-class="completeDia">
            <img src="@/assets/images/alert-icon.png" />
            <h1>{{$t('uploadFile.COMPLETED')}}!</h1>
            <h3>{{$t('uploadFile.SUCCESS')}}</h3>
            <a :href="'https://mumbai.polygonscan.com/tx/'+txHash" target="_blank">{{txHash}}</a>
            <a class="a-close" @click="finishClose">{{$t('uploadFile.CLOSE')}}</a>
        </el-dialog>

        <el-dialog title="" :visible.sync="failTransaction" :width="width"
            custom-class="completeDia">
            <img src="@/assets/images/error.png" />
            <h1>{{$t('uploadFile.Fail')}}!</h1>
            <h3>{{$t('uploadFile.FailTIP')}}</h3>
            <a :href="'https://mumbai.polygonscan.com/tx/'+txHash" target="_blank">{{txHash}}</a>
            <a class="a-close" @click="failTransaction=false">{{$t('uploadFile.CLOSE')}}</a>
        </el-dialog>

        <el-dialog
        :visible.sync="fileUploadVisible" :show-close="false" :close-on-click-modal="false"
        :width="widthUpload" custom-class="fileUpload">
        <span slot="title">{{$t('uploadFile.File_uploading')}}... {{percentIn?'('+percentIn+')':''}}</span>
        <h3>{{$t('uploadFile.File_uploading_tooltip')}}</h3>
        <img src="@/assets/images/upload.gif" class="gif_img" alt="">
        </el-dialog>

        <el-dialog title="" :visible.sync="paymentPopup" :width="width"
            custom-class="completeDia">
            <img src="@/assets/images/box-important.png" />
            <h2>{{$t('uploadFile.file_uploaded')}}</h2>
            <h4>{{$t('uploadFile.file_uploaded_tip')}}</h4>
            <h4>{{$t('uploadFile.file_uploaded_tip01')}}</h4>
            <a class="a-close" @click="paymentPopup=false">{{$t('uploadFile.OK')}}</a>
        </el-dialog>

        <el-dialog title="" :visible.sync="paymentPopup01" :width="width"
            custom-class="completeDia">
            <img src="@/assets/images/box-important.png" />
            <h2>{{$t('uploadFile.file_uploaded')}}</h2>
            <h4>{{$t('uploadFile.file_uploaded_tip02')}}</h4>
            <h4>{{$t('uploadFile.file_uploaded_tip01')}}</h4>
            <a class="a-close" @click="paymentPopup01=false">{{$t('uploadFile.OK')}}</a>
        </el-dialog>

        <el-dialog title="" :visible.sync="metamaskLoginTip" :width="width"
            custom-class="completeDia">
            <img src="@/assets/images/box-important.png" />
            <h4>{{$t('fs3Login.toptip_01')}} {{network.name}} {{$t('fs3Login.toptip_02')}} <b>{{$t('fs3Login.toptip_Network')}}</b>.</h4>
            <a class="a-close" @click="metamaskLoginTip=false">{{$t('uploadFile.OK')}}</a>
        </el-dialog>
    </div>
</template>

<script>
    import bus from "@/components/bus";
    import * as myAjax from "@/api/uploadFile";
    import NCWeb3 from "@/utils/web3";
    import axios from 'axios'
    import first_contract_json from "@/utils/swanPayment.json";
    import erc20_contract_json from "@/utils/ERC20.json";
    let contract_erc20
    let that
    export default {
        name: "uploadFiles",
        data() {
            var validateDuration = (rule, value, callback) => {
                if (!value) {
                    return callback(new Error(that.$t('uploadFile.Duration_tip')));
                }
                setTimeout(() => {
                    if (value < 180) {
                        callback(new Error(' '));
                        callback(that.calculation(1));
                    }else if (value > 540) {
                        callback(new Error(' '));
                        callback(that.calculation(2));
                    } else {
                        callback(that.calculation());
                    }
                }, 100);
            };
            return {
                tableData: {
                    file_name: '',
                    task_name: '',
                    payload_cid: ''
                },
                ruleForm: {
                    duration: '365',
                    fileList: [],
                    file_size: '',
                    file_size_byte: '',
                    storage_cost: '',
                    lock_plan: '',
                    lock_plan_tip: false,
                    amount: '',
                    amount_minprice: 0,
                    gaslimit: this.$root.PAY_GAS_LIMIT,
                },
                rules: {
                    duration: [
                        { validator: validateDuration, trigger: 'blur'}
                    ],
                    fileList: [
                        { type: 'array', required: true, message: this.$t('uploadFile.file_name_tip'), trigger: 'change' }
                    ]
                },
                loading: false,
                bodyWidth: document.documentElement.clientWidth < 1024 ? true : false,
                fileListTip: false,
                storage: 0,
                biling_price: 0,
                storage_cost_low: 0,
                storage_cost_average: 0,
                storage_cost_high: 0,
                center_fail: false,
                centerDialogVisible: false,
                modelClose: true,
                width: document.body.clientWidth>600?'400px':'95%',
                widthUpload: document.body.clientWidth>600?'450px':'95%',
                gatewayContractAddress: this.$root.SWAN_PAYMENT_CONTRACT_ADDRESS,
                recipientAddress: this.$root.RECIPIENT,
                usdcAddress: this.$root.USDC_ADDRESS,
                hashload: false,
                timer: '',
                usdcAvailable: '',
                network: {
                    name: '',
                    unit: '',
                    text: false
                },
                finishTransaction: false,
                failTransaction: false,
                txHash: '',
                fileUploadVisible: false,
                paymentPopup: false,
                paymentPopup01: false,
                percentIn: '',
                loadMetamaskPay: false,
                metamaskLoginTip: false
            };
        },
        components: {},
        watch: {
            'ruleForm.fileList': function(){
                this.ruleForm.lock_plan = ''
                this.$refs['ruleForm'].validate((valid) => {});
                if(this.ruleForm.file_size_byte > 0){
                    this.calculation()
                }
            }
        },
        methods: {
            calculation(type){
                if(type && type == 2){
                    this.ruleForm.duration = '540'
                }else if(type && type == 1){
                    this.ruleForm.duration = '180'
                }
                this.ruleForm.storage_cost = this.ruleForm.file_size_byte * this.ruleForm.duration * this.storage / 365
                this.ruleForm.amount_minprice = Number(this.ruleForm.storage_cost * this.biling_price).toFixed(9)
                this.storage_cost_low = Number(this.ruleForm.storage_cost * this.biling_price * 2).toFixed(9)
                this.storage_cost_average = Number(this.ruleForm.storage_cost * this.biling_price * 3).toFixed(9)
                this.storage_cost_high = Number(this.ruleForm.storage_cost * this.biling_price * 5).toFixed(9)
            },
            agreeChange(val){
                this.ruleForm.lock_plan_tip = false
                switch (val) {
                    case '1':
                        this.ruleForm.amount = this.storage_cost_low
                        break;
                    case '2':
                        this.ruleForm.amount = this.storage_cost_average
                        break;
                    case '3':
                        this.ruleForm.amount = this.storage_cost_high
                        break;
                    default:
                        this.ruleForm.amount = this.storage_cost_low
                        return;
                }
            },
            submitForm(formName) {
                let _this = this;
                _this.$refs[formName].validate((valid) => {
                    if (valid) {
                        if(_this.metaAddress&&_this.networkID!=80001) {
                            _this.metamaskLoginTip = true
                            return false
                        }
                        if(!_this.ruleForm.lock_plan || _this.ruleForm.amount<=0){
                            _this.ruleForm.lock_plan_tip = true
                            return false
                        }

                        if(_this.metaAddress){
                            // 授权代币
                            contract_erc20 = new web3.eth.Contract( erc20_contract_json );
                            contract_erc20.options.address = _this.usdcAddress
                            // 查询剩余代币余额为：
                            contract_erc20.methods.balanceOf(_this.metaAddress).call()
                            .then(resultUSDC => {
                                _this.usdcAvailable = web3.utils.fromWei(resultUSDC, 'ether');
                                console.log('Available:', _this.usdcAvailable)
                                // console.log(_this.ruleForm.amount, _this.usdcAvailable)

                                // 判断支付金额是否大于代币余额
                                if(_this.ruleForm.amount > _this.usdcAvailable ){
                                    _this.$message.error('Insufficient balance')
                                    return false
                                }

                                // 通过 FormData 对象上传文件
                                var formData = new FormData()
                                formData.append('file', _this._file)
                                formData.append('duration', _this.ruleForm.duration)
                                formData.append('wallet_address', _this.metaAddress)
                                // formData.append('task_name', _this.ruleForm.task_name)
                                _this.loading = true
                                _this.fileUploadVisible = true


                                let xhr = new XMLHttpRequest()
                                xhr.open("POST", `${process.env.BASE_PAYMENT_GATEWAY_API}api/v1/storage/ipfs/upload`, true);   // 设置xhr得请求方式和url。
                                xhr.withCredentials = false
                                const token = _this.$store.getters.accessToken
                                if (token) {
                                    // xhr.setRequestHeader(
                                    // "Authorization",
                                    // "Bearer " + _this.$store.getters.accessToken
                                    // )
                                }
                                let i = 0;

                                xhr.onreadystatechange = function() {   // 等待ajax请求完成。
                                    if (xhr.status === 200) { 
                                        if(xhr.responseText){
                                            _this.fileUploadVisible = false
                                            let res = JSON.parse(xhr.responseText)
                                            i += 1
                                            if(i <= 1){
                                                if(res.status == "success"){
                                                    if(res.data.need_pay == 0 || res.data.need_pay == 2 || res.data.need_pay == 4){
                                                        contract_erc20.methods.allowance(_this.gatewayContractAddress, _this.metaAddress).call()
                                                        .then(resultUSDC => {
                                                            console.log('allowance：'+ resultUSDC);
                                                            if(resultUSDC < web3.utils.toWei(_this.ruleForm.amount, 'ether')){
                                                                contract_erc20.methods.approve(_this.gatewayContractAddress, web3.utils.toWei(_this.ruleForm.amount, 'ether')).send({from:  _this.metaAddress})
                                                                .then(receipt => {
                                                                    // console.log(receipt)
                                                                })
                                                                .catch(error => {
                                                                    // console.log('errorerrorerror', error)
                                                                })
                                                            }
                                                            _this.contractSend(res.data.payload_cid)
                                                        })
                                                    }else if(res.data.need_pay == 3){
                                                        _this.paymentPopup01 = true
                                                        _this.loading = false
                                                        return false
                                                    }else{
                                                        // res.data.need_pay == 1 && 其他情况
                                                        _this.paymentPopup = true
                                                        _this.loading = false
                                                        return false
                                                    }
                                                }else{
                                                    _this.$message.error('Fail')
                                                }
                                            }
                                        }
                                    } else {
                                        _this.loading = false
                                        _this.fileUploadVisible = false
                                    }
                                    xhr.upload.addEventListener("error", event => {
                                        _this.$message.error('Fail')
                                    })

                                    xhr.upload.addEventListener("progress", event => {
                                        if (event.lengthComputable) {
                                            let loaded = event.loaded
                                            let total = event.total
                                            console.log('total-loaded', total, loaded)
                                            let percentIn = Math.floor(event.loaded / event.total * 100);
                                            // 设置进度显示
                                            _this.percentIn = percentIn+'%'
                                            console.log(percentIn+'%-')
                                        }
                                    })
                                };
                                // 获取上传进度
                                xhr.upload.onprogress = function(event) { 
                                    console.log('event.loaded', event.loaded)
                                    console.log('event.total', event.total)
                                    if (event.lengthComputable) {
                                        let percentIn = Math.floor(event.loaded / event.total * 100);
                                        // 设置进度显示
                                        _this.percentIn = percentIn+'%'
                                        console.log(percentIn+'%')
                                    }
                                };
                                xhr.send(formData);
                                return false

                                // 发起请求
                                // axios.post(`${process.env.BASE_PAYMENT_GATEWAY_API}api/v1/storage/ipfs/upload `, formData,{
                                //     headers: {
                                //     'Authorization': "Bearer "+_this.$store.getters.accessToken
                                //     },
                                // })
                                // .then((res) => {
                                //     // console.log('_RequestUploads_', res)
                                //     _this.fileUploadVisible = false
                                //     if (res.data.status == "success") {
                                //         if(!res.data.data.need_pay){
                                //             _this.paymentPopup = true
                                //             _this.loading = false
                                //             return false
                                //         }
                                //         contract_erc20.methods.allowance(_this.gatewayContractAddress, _this.metaAddress).call()
                                //         .then(resultUSDC => {
                                //             console.log('allowance：'+ resultUSDC);
                                //             if(resultUSDC < web3.utils.toWei(_this.ruleForm.amount, 'ether')){
                                //                 contract_erc20.methods.approve(_this.gatewayContractAddress, web3.utils.toWei(_this.ruleForm.amount, 'ether')).send({from:  _this.metaAddress})
                                //                 .then(receipt => {
                                //                     // console.log(receipt)
                                //                 })
                                //             }
                                //             _this.contractSend(res.data.data.payload_cid)
                                //         })
                                //         // _this.$router.push({name: 'my_files'})
                                //     } else {
                                //         _this.$message.error('Fail')
                                //     }
                                // }).catch(error => {
                                //     console.log(error)
                                //     _this.loading = false
                                //     _this.fileUploadVisible = false
                                // })
                            })
                        }
                    } else {
                        console.log('error submit!!');
                        return false;
                    }
                });
            },
            contractSend(cid){
                let _this = this
                // 合约转账
                let contract_instance = new web3.eth.Contract( first_contract_json );
                contract_instance.options.address = _this.gatewayContractAddress
                // console.log( 'contract_instance合约实例：', contract_instance );
                // console.log(contract_instance.options.jsonInterface)

                let payObject = {
                    from: _this.metaAddress,
                    gas: web3.utils.toHex(_this.ruleForm.gaslimit),
                    // gasPrice: web3.utils.toHex(web3.utils.toWei(_this.ruleForm.gasprice + '', 'gwei')),
                    // value: web3.utils.toHex(web3.utils.toWei(_this.ruleForm.amount, 'ether')),
                };
                
                let lockObj = {
                    id: cid,
                    minPayment: web3.utils.toWei(String(_this.ruleForm.amount_minprice), 'ether'),
                    amount: web3.utils.toWei(_this.ruleForm.amount, 'ether'),
                    lockTime: 86400 * Number(_this.$root.LOCK_TIME), // one day
                    recipient: _this.recipientAddress, //todo:
                }
                
                contract_instance.methods.lockTokenPayment(lockObj)
                .send(payObject)
                .on('transactionHash', function(hash){
                    // console.log('hash console:', hash);
                    _this.loadMetamaskPay = true
                    _this.loading = false
                    _this.txHash = hash
                })
                .on('confirmation', function(confirmationNumber, receipt){
                    // console.log('confirmationNumber console:', confirmationNumber, receipt);
                })
                .on('receipt', function(receipt){
                    // receipt example
                    // console.log('receipt console:', receipt);
                    _this.checkTransaction(receipt.transactionHash, cid)
                    _this.txHash = receipt.transactionHash
                })
                .on('error', function(error){
                    // console.log('error console:', error)
                    // console.error
                    _this.loading = false
                    _this.loadMetamaskPay = false
                    _this.failTransaction = true
                }); 
            },
            checkTransaction(txHash, cid) {
                let _this = this
                web3.eth.getTransactionReceipt(txHash).then(
                    res => {
                        console.log('checking ... ');
                        if (!res) { return _this.timer = setTimeout(() => { _this.checkTransaction(txHash, cid); }, 2000); }
                        else {
                            setTimeout(function(){
                                _this.loading = false
                                _this.loadMetamaskPay = false
                                clearTimeout(_this.timer)
                                _this.finishTransaction = true
                            }, 2000)
                        }
                    },
                    err => { console.error(err); }
                );
            },
            finishClose(){
                this.finishTransaction = false
                this.$router.push({name: 'my_files'})
            },
            // 文件上传
            uploadFile(params) {
                this._file = params.file;
                const isLt2M = this._file.size / 1000 / 1000 < 2;  // or 1024
                this.ruleForm.file_size = this.sizeChange(this._file.size)
                this.ruleForm.file_size_byte = this.byteChange(this._file.size)
                console.log('bytes', this._file.size)
                if (!isLt2M) {
                    // this.$message.error(this.$t('deal.upload_form_file_tip'))
                    this.fileListTip = true
                    return false
                }else{
                    this.fileListTip = false
                }
            },
            sizeChange(bytes){
                if (!bytes) return "-";
                if (bytes === 0) return '0 B';
                var k = 1000, // or 1024
                    sizes = ['bytes', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB'],
                    i = Math.floor(Math.log(bytes) / Math.log(k));
                return (bytes / Math.pow(k, i)).toPrecision(3) + ' ' + sizes[i];
            },
            byteChange(limit){
                var size = "";
                // 只转换成GB
                if(limit <= 0){
                    return '-'
                }else{
                    size = limit/( 1000 * 1000 * 1000)  //or 1024
                }
                return size
                // return Number(size).toFixed(3);
            },
            handleChange(file, fileList) {
                if (fileList.length > 0) {
                    this.ruleForm.fileList = [fileList[fileList.length - 1]]  // 这一步，是 展示最后一次选择的csv文件
                }
            },
            handleRemove(file, fileList) {
                // console.log(file, fileList);
                this.ruleForm.fileList = []
                this.ruleForm.file_size = ''
                this.ruleForm.file_size_byte = ''
            },
            stats(){
                let _this = this
                _this.loading = true
                if(_this.$root.SWAN_PAYMENT_CONTRACT_ADDRESS){
                    _this.gatewayContractAddress = _this.$root.SWAN_PAYMENT_CONTRACT_ADDRESS
                    _this.usdcAddress = _this.$root.USDC_ADDRESS
                    _this.recipientAddress = _this.$root.RECIPIENT
                    let stats_api = `${process.env.BASE_API}stats/storage?wallet_address=${_this.metaAddress}`
                    axios.get(stats_api, {
                        headers: {
                            // 'Authorization': "Bearer "+ _this.$store.getters.accessToken
                        },
                    }).then(res => {
                        if(res.data.data){
                            let cost = res.data.data.average_price_per_GB_per_year.split(" ")
                            if(cost[0]) _this.storage = cost[0]
                        }
                        setTimeout(function(){
                            _this.loading = false
                        }, 2000)
                    }).catch(error => {
                        console.log(error)
                        _this.loading = false
                    })
                    
                    let billing_api = `${process.env.BASE_PAYMENT_GATEWAY_API}api/v1/billing/price/filecoin?wallet_address=${_this.metaAddress}`
                    axios.get(billing_api, {
                        headers: {
                            // 'Authorization': "Bearer "+ _this.$store.getters.accessToken
                        },
                    }).then(res => {
                        if(res.data.data){
                            _this.biling_price = res.data.data
                        }
                    }).catch(error => {
                        console.log(error)
                    })
                }else {
                    setTimeout(function(){
                        _this.stats()
                    }, 1000)
                }
            }
        },
        mounted() {
            let _this = this
            that = _this
            _this.stats()
            _this.$store.dispatch("setRouterMenu", 0);
            _this.$store.dispatch('setHeadertitle', _this.$t('navbar.Upload_files'))
            document.onkeydown = function(e) {
                if (e.keyCode === 13) {
                }
            }
        },
        filters: {
            NumStorage(value) {
                if (!value) return "-";
                return value.toFixed(8);
            },
        },
        computed: {
            metaAddress() {
                return this.$store.getters.metaAddress
            },
            networkID() {
                return this.$store.getters.networkID
            }
        },
    };
</script>


<style scoped lang="scss">
    .el-dialog__wrapper /deep/ {  
        display: flex;
        align-items: center;
        .metaM{
            .el-dialog__body{
                padding: 0.25rem 0.25rem 0.2rem;
                .el-row{
                border-radius: 0.08rem;
                padding: 0.16rem;
                margin: 0.12rem 0px;
                border: 1px solid rgb(240, 185, 11);
                text-align: center;
                display: flex;
                -webkit-box-pack: justify;
                justify-content: space-between;
                -webkit-box-align: center;
                align-items: center;
                transition: all 0.3s ease 0s;
                background: rgba(240, 185, 11, 0.1);
                position: static;
                .el-col{
                    text-align: left;
                    font-size: 0.14rem;
                    img{
                    float: right;
                    height: 0.24rem;
                    }
                }
                }
            }
            .el-dialog__footer{
                padding: 0 0.25rem 0.3rem;
                .dialog-footer{
                .el-button{
                    width: 100%;
                    font-size: 0.14rem;
                    height: 0.4rem;
                    padding: 0;
                    background: #5c3cd3;
                    color: #fff;
                    border-radius: 0.08rem;
                    &:hover{
                    background: #4326ab;
                    }
                }
                p{
                    font-size: 0.13rem;
                    line-height: 1.5;
                    color: red;
                    margin: 0.1rem 0 0;
                }
                }
            }
        }
        .completeDia{
            text-align: center;
            .el-dialog__header{
            display: none;
            }
            img{
                display: block;
                max-width: 100px;
                margin: auto;
            }
            h1{
                margin: 0.32rem auto 0.1rem;
                font-size: 0.32rem;
                font-weight: 500;
                line-height: 1.2;
                color: #191919;
                word-break: break-word;
            }
            h2{
                margin: 0.1rem auto 0.1rem;
                font-size: 19px;
                font-weight: 600;
                line-height: 1.2;
                color: #191919;
                word-break: break-word;
                text-align: center;
            }
            h4{
                font-size: 14px;
                font-weight: 500;
                line-height: 1.2;
                color: #191919;
                word-break: break-word;
                text-align: center;
            }
            h3, a{
                font-size: 0.16rem;
                font-weight: 500;
                line-height: 1.2;
                color: #191919;
                word-break: break-word;
            }
            a{
                text-decoration: underline;
                color: #007bff;
            }
            a.a-close{
                padding: 5px 45px;
                background: #5c3cd3;
                color: #fff;
                border-radius: 10px;
                cursor: pointer;
                margin: 0.2rem auto 0;
                display: block;
                width: max-content;
                text-decoration: unset;
            }
        }
        .fileUpload{
            .el-dialog__header{
                font-size: 0.16rem;
                font-weight: 600;
            }
            .el-dialog__body{
                padding: 0.1rem 0.2rem 0.2rem;
                h3{   
                    margin: 0 0 0.1rem;
                    font-size: 0.14rem;
                    font-weight: normal;
                    line-height: 1.2;
                    color: #666;
                    word-break: break-word;
                }
                .gif_img{
                    max-width: 200px;
                    width: 100%;
                    margin: auto;
                    display: block;
                }
            }
        }
    }
    #Create {
        position: relative;
        height: calc(100% - 0.6rem);
        padding: 0.4rem 0.2rem 0;
        font-size: 0.24rem;
        .loadMetamaskPay{
            position: absolute;
            left: 0;
            top: 0;
            right: 0;
            bottom: 0;
            z-index: 99;
            display: flex;
            align-items: center;
            justify-content: center;
            background-color: rgba(255,255,255,.9);
            .el-loading-spinner{
                top: 0;
                position: relative;
                margin: 0 0 0.2rem;
            }
            p{
                font-size: 14px;
                font-weight: 600;
                color: #666;
            }
        }
        .el-alert /deep/{
            position: absolute;
            left: 0;
            top: 0;
            .el-alert__content{
                display: flex;
                align-items: center;
            }
        }
        .upload{
            padding: 0 0.17rem 0.4rem;
            margin: 0 auto 0.2rem;
            background-color: #fff;
            border-radius: 0.1rem;
            overflow: hidden;
            .title{
                width: 100%;
                margin: 0 0 0.1rem;
                text-align: left;
                font-size: 0.1972rem;
                color: #000;
                line-height: 0.42rem;
                text-indent: 0.08rem;
            }
            .upload_title{
                width: 100%;
                margin: 0.1rem 0 0.3rem;
                text-align: left;
                font-size: 0.15rem;
                font-weight: 600;
                color: #000;
                line-height: 1.5;
                text-indent: 0;
            }
            .upload_form{
                // display: flex;
                // align-items: baseline;
                width: 80%; 
                margin: auto; 
                justify-content: flex-start;
                .el-form /deep/{
                    width: 96%;
                    margin: 0 2%;
                    .el-form-item{
                        display: flex;
                        align-items: center;
                        width: auto;
                        margin: 0.15rem auto;
                        @media screen and (max-width:441px){ 
                            flex-wrap: wrap;
                        }
                        .el-form-item__label{
                            display: flex;
                            justify-content: flex-end;
                            align-items: center;
                            width: 45%;
                            padding: 0 2% 0 0;
                            // max-width: 2rem;
                            line-height: 1.5;
                            text-align: left;
                            font-size: 0.1372rem;
                            white-space: normal;
                            color: #000;
                            font-weight: 500;
                            text-shadow: 0 0 black;
                            text-align: right;
                            img{
                                width: 0.16rem;
                                height: 0.16rem;
                                margin: 0 0 0 5px;
                                cursor: pointer;
                            }
                            &::before{
                                display: none;
                            }
                            @media screen and (max-width:441px){ 
                                width: 100%;
                                padding-bottom: 0.05rem;
                                justify-content: center;
                            }
                        }
                        .el-form-item__content{
                            display: flex;
                            align-items: center;
                            width: 53%;
                            font-size: 0.14rem;
                            white-space: normal;
                            word-break: break-word;
                            line-height: 1.5;
                            color: #666;
                            @media screen and (max-width:1024px){
                                flex-wrap: wrap;
                            }
                            @media screen and (max-width:768px){
                                font-size: 14px;
                            }
                            @media screen and (max-width:441px){ 
                                flex-wrap: wrap;
                                width: 100%;
                                justify-content: center;
                            }
                            small{
                                margin: 2px 0 0 5px;
                                font-size: 12px;
                                @media screen and (max-width:1024px){
                                    width: 100%;
                                }
                                @media screen and (max-width:441px){
                                    text-align: center;
                                }
                            }
                            h4{
                                width: 100%;
                                font-size: inherit;
                                font-weight: 500;
                                line-height: 1.7;
                            }
                            h5{
                                width: 90%;
                                margin-top: 5px;
                                font-size: 0.11rem;
                                font-weight: 500;
                                line-height: 1.2;
                                color: #737373;
                                @media screen and (max-width:600px){
                                    font-size: 12px;
                                }
                            }
                            .el-tag, .el-button--small{
                                margin: 0 5px 5px 0;
                            }
                            .el-input{
                                width: auto;
                                .el-input__inner{
                                    height: 35px;
                                    font-size: inherit;
                                    line-height: 34px;
                                }
                                .el-input__suffix{
                                    display: none;
                                }
                            }
                            .el-input-number{
                                max-width: 140px;
                                line-height: 34px;
                            }
                            .el-input-number.is-controls-right .el-input-number__decrease, 
                            .el-input-number.is-controls-right .el-input-number__increase{
                                line-height: 17px;
                            }
                            .el-form-item__error {
                                padding-top: 0;
                                margin: 0 0.1rem;
                                position: relative;
                                float: right;
                            }
                            .el-textarea{
                                width: 90% !important;
                            }
                            .upload-demo{
                                display:flex;
                                align-items: center;
                                flex-wrap: wrap;
                                .el-upload-list__item:first-child{
                                    margin-top: 0;
                                }
                                .el-upload--text{
                                    float: left;
                                    width: auto;
                                    height: auto;
                                    text-align: left;
                                    border: 0;
                                    .el-button--primary{
                                        height: 0.32rem;
                                        padding: 0 0.2rem;
                                        margin: 0 5px 0 0;
                                        line-height: 0.32rem;
                                        background-color:transparent;
                                        border: 1px solid #2c4c9e;
                                        border-radius: 0.08rem;
                                        color: #2c4c9e;
                                        font-size: 0.1372rem;
                                    }
                                }
                                .el-upload-list{
                                    width: 100%;
                                    float: none;
                                    clear: both;
                                }
                            }
                            .el-upload__tip{
                                // float: left;
                                height: 100%;
                                align-items: center;
                                display: flex;
                                margin: 0 0 0 0.1rem;
                                color: #737373;
                                line-height: 1;
                                font-size: 0.12rem;
                                @media screen and (max-width:600px){
                                    font-size: 14px;
                                }
                            }
                            .el-radio{
                                .el-radio__inner{
                                    border-color: #d9d9d9;
                                    background-color: #d9d9d9;
                                }
                            }
                            .el-radio.is-checked{
                                .el-radio__inner{
                                    border-color: #0b318f;
                                    background-color: #0b318f;
                                }
                                .el-radio__inner::after{
                                    width: 6px;
                                    height: 6px;
                                }
                            }
                        }
                    }
                }
            }
            .upload_plan{
                width: 80%;
                margin: auto; 
                justify-content: flex-start;
                .title{
                    display: flex;
                    align-items: center;
                    justify-content: center;
                    margin: 0.3rem 0 0;
                    line-height: 1.5;
                    text-align: center;
                    font-size: 0.15rem;
                    white-space: normal;
                    color: #000;
                    font-weight: 500;
                    text-shadow: 0 0 black;
                    text-indent: 0;
                    img{
                        width: 0.16rem;
                        height: 0.16rem;
                        margin: 0 0 0 5px;
                        cursor: pointer;
                    }
                }
                .desc{
                    margin: 0 0 0.1rem;
                    line-height: 1.5;
                    text-align: center;
                    font-size: 0.13rem;
                    white-space: normal;
                    color: #999;
                    font-weight: normal;
                }
                .upload_plan_radio{
                    .el-radio-group /deep/{
                        width: 100%;
                        display: flex;
                        justify-content: center;
                        align-items: center;
                        .el-radio{
                            min-width: 20%;
                            height: auto;
                            padding: 0 0.1rem 0.15rem;
                            // line-height:30px;
                            .el-radio__input{
                                display: none;
                            }
                            .el-radio__label{
                                .title{
                                    margin: 0 0 0.05rem;
                                    font-size: 0.13rem;
                                }
                                .cont{
                                    font-size: 0.145rem;
                                    font-weight: bold;
                                    line-height: 1.5;
                                    text-align: center;
                                }
                            }
                        }
                        .el-radio:nth-child(3n+1){
                            .el-radio__label{
                                .cont{
                                    color: #56c4a6;
                                }
                            }
                        }
                        .el-radio:nth-child(3n+2){
                            .el-radio__label{
                                .cont{
                                    color: #4a92d3;
                                }
                            }
                        }
                        .el-radio:nth-child(3n+3){
                            .el-radio__label{
                                .cont{
                                    color: #922b26;
                                }
                            }
                        }
                        .is-checked{
                            position: relative;
                            &:after {
                                content: "";
                                display: block;
                                height: 25px;
                                width: 25px;
                                background-image: url(../assets/images/plan.png);
                                background-size: 100%;
                                position: absolute;
                                right:0;
                                top:0;
                            }
                        }
                        .el-radio:hover{
                            background-color: rgba(64,158,255,0.1);
                        }

                    }

                }
            }
            .upload_bot{
                display: flex;
                justify-content: center;
                align-items: center;
                width: 100%;
                margin: 0.25rem auto 0.15rem;
                .el-button /deep/{
                    height: 0.35rem;
                    padding: 0 0.4rem;
                    margin-left: 0;
                    background-color: #0b318f;
                    line-height: 0.35rem;
                    font-size: 0.1372rem;
                    color: #fff;
                    border: 0;
                }
            }
            .upload_result /deep/{
                width: 75%; 
                margin: 0.6rem auto 0.2rem; 
                .el-col{
                    display: flex;
                    align-items: flex-start;
                    font-size: 0.12rem;
                    color: #222;
                    margin: 0.15rem 0 0;
                    h5{
                        width: 100%;
                        font-size: 0.14rem;
                        font-weight: 500;
                        line-height: 1.3;
                        color: #000;
                    }
                    label{
                        width: 100px;
                        margin: 0 0 0 0.2rem;
                    }
                    p{
                        word-break: break-word;
                    }
                }
            }
        }


        /deep/ .el-list-enter-active,
        /deep/ .el-list-leave-active {
            transition: none;
        }

        /deep/ .el-list-enter,
        /deep/ .el-list-leave-active {
            opacity: 0;
        }
    }
    // #Create::-webkit-scrollbar-track {
    //     background: #ccc;
    // }
    // #Create::-webkit-scrollbar {
    //     width: 6px;
    //     background: #ccc;
    // }
    // #Create::-webkit-scrollbar-thumb {
    //     background: #f2f2f2;
    //     border-radius: 0.08rem;
    // }
    @media screen and (max-width: 999px) {
        #Create {
            padding: 0.1rem;
            .upload{
                padding: 0.1rem;
                .upload_form{
                    width: 90%;
                    flex-wrap: wrap;
                    .el-form /deep/{
                        width: 95%;
                    }
                }
            }
        }
    }
    @media screen and (max-width: 600px){
        #Create {
            .upload{
                 .upload_plan {
                     .upload_plan_radio {
                         .el-radio-group /deep/{
                             .el-radio{
                                 width: 32%;
                                 margin: auto;
                             }
                         }
                     }
                 }
            }
        }
    }
    @media screen and (max-width: 441px){
        #Create {
            .upload{
                 .title{
                    text-indent: 0;
                    font-size: 0.19rem;
                    line-height: 1.5;
                 }
                 .upload_form {
                     width: 100%;
                 }
                 .upload_plan {
                     width: 95%;
                 }
                 .upload_result{
                     width: 90%;
                     margin: 0.2rem auto 0;
                     .el-col{
                         flex-wrap: wrap;
                         label{
                            margin: 0 0 5px;
                            width: 100%;
                            font-weight: 600;
                         }
                     }
                 }
            }
        }
    }
</style>
