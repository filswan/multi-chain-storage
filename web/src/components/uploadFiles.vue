<template>
    <div id="Create">
        <el-dialog :modal="true" :close-on-click-modal="false" :width="widthDia" :visible.sync="uploadDigShow"
            custom-class="uploadDig"
            :before-close="closeDia">
            <div class="loadMetamaskPay" v-if="loading">
                <div>
                    <div class="el-loading-spinner"><svg viewBox="25 25 50 50" class="circular"><circle cx="50" cy="50" r="20" fill="none" class="path"></circle></svg><!----></div>
                    <p v-if="loadMetamaskPay">{{$t('uploadFile.payment_tip')}}</p>
                </div>
            </div>
            <template slot="title">
                {{$t('uploadFile.upload')}}  
                
                <el-tooltip effect="dark" :content="$t('uploadFile.uploadFile_title')" placement="top">
                    <img src="@/assets/images/info.png"/>
                </el-tooltip>
            </template>
            <div class="upload_form">
                <el-form :model="ruleForm" status-icon :rules="rules" ref="ruleForm" class="demo-ruleForm">
                    <el-form-item prop="fileList" :label="$t('uploadFile.upload')" :style="{'align-items': ruleForm.fileList_tip?'flex-start':'center'}">
                        <div>
                            <el-upload
                                class="upload-demo"
                                ref="uploadFileRef"
                                action="customize"
                                :http-request="uploadFile"
                                :file-list="ruleForm.fileList"
                                :on-change="handleChange"
                                :on-remove="handleRemove">
                                <el-button size="small" type="primary">
                                    <img src="@/assets/images/my_file/icon_Upload@2x.png" alt="">
                                    {{$t('uploadFile.upload')}}
                                </el-button>
                            </el-upload>
                            <p v-if="ruleForm.fileList.length>0" style="display: flex;align-items: center;">
                                <svg t="1637031488880" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="3310" style="width: 14px;height: 14px;margin: 0 6px 0 5px;"><path d="M512 1024a512 512 0 1 1 512-512 32 32 0 0 1-32 32h-448v448a32 32 0 0 1-32 32zM512 64a448 448 0 0 0-32 896V512a32 32 0 0 1 32-32h448A448 448 0 0 0 512 64z" fill="#999999" p-id="3311"></path><path d="M858.88 976a32 32 0 0 1-32-32V640a32 32 0 0 1 32-32 32 32 0 0 1 32 32v304a32 32 0 0 1-32 32z" fill="#999999" p-id="3312"></path><path d="M757.12 773.12a34.56 34.56 0 0 1-22.4-8.96 32 32 0 0 1 0-45.44l101.12-101.12a32 32 0 0 1 45.44 0 30.72 30.72 0 0 1 0 44.8l-101.12 101.76a34.56 34.56 0 0 1-23.04 8.96z" fill="#999999" p-id="3313"></path><path d="M960 773.12a32 32 0 0 1-22.4-8.96l-101.76-101.76a32 32 0 0 1 0-44.8 32 32 0 0 1 45.44 0l101.12 101.12a32 32 0 0 1-22.4 54.4z" fill="#999999" p-id="3314"></path></svg>
                                {{ruleForm.file_size}}
                            </p>
                            <p v-if="ruleForm.fileList_tip" style="color: #F56C6C;font-size: 12px;line-height: 1;">{{ruleForm.fileList_tip_text}}</p>
                        </div>
                    </el-form-item>
                    <el-form-item prop="duration">
                        <template slot="label">
                            {{$t('uploadFile.Duration')}} 
                            
                            <el-tooltip effect="dark" :content="$t('uploadFile.Duration_tooltip')" placement="top">
                                <img src="@/assets/images/info.png"/>
                            </el-tooltip>
                        </template>
                        {{ruleForm.duration}} {{$t('components.day')}}
                    </el-form-item>
                    <el-form-item prop="storage_copy">
                        <template slot="label">
                            {{$t('uploadFile.Storage_copy')}} 
                            
                            <el-tooltip effect="dark" :content="$t('uploadFile.Storage_copy_tooltip')" placement="top">
                                <img src="@/assets/images/info.png"/>
                            </el-tooltip>
                        </template>
                        {{ruleForm.storage_copy}}
                    </el-form-item>
                    <el-form-item prop="Free_Storage_Capacity">
                        <template slot="label">
                            {{$t('uploadFile.Free_Storage_Capacity')}} 
                            
                            <el-tooltip effect="dark" :content="$t('uploadFile.Free_Storage_Capacity_tooltip')" placement="top">
                                <img src="@/assets/images/info.png"/>
                            </el-tooltip>
                        </template>
                        <span  style="color:#2C7FF8">{{free_quota_per_month-free_usage | byteStorage}} GB</span> 
                    </el-form-item>
                    <el-form-item prop="storage_cost">
                        <template slot="label">
                            {{$t('uploadFile.Estimated_Storage_Cost')}} 
                            
                            <el-tooltip effect="dark" :content="$t('uploadFile.Estimated_Storage_Cost_tooltip')" placement="top">
                                <img src="@/assets/images/info.png"/>
                            </el-tooltip>
                        </template>
                        <span v-if="free" style="color:#2C7FF8;text-transform: uppercase;">{{$t('uploadFile.filter_status_Free')}}</span>
                        <span v-else style="color:#2C7FF8">{{ruleForm.storage_cost | NumStorage}} FIL</span> 
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
                    <div class="upload_plan_radio" :class="{'upload_plan_radio_free': free}">
                        <el-radio-group :disabled="free?true:false" v-model="ruleForm.lock_plan" @change="agreeChange">
                            <el-radio label="1" border>
                                <div class="title">{{$t('uploadFile.Low')}}</div>
                                <div class="cont">
                                    {{free?0:storage_cost_low}} USDC
                                </div>
                            </el-radio>
                            <el-radio label="2" border>
                                <div class="title">{{$t('uploadFile.Average')}}</div>
                                <div class="cont">
                                    {{free?0:storage_cost_average}} USDC
                                </div>
                            </el-radio>
                            <el-radio label="3" border>
                                <div class="title">{{$t('uploadFile.High')}}</div>
                                <div class="cont">
                                    {{free?0:storage_cost_high}} USDC
                                </div>
                            </el-radio>
                        </el-radio-group>
                    </div>
                </div>
                <div class="upload_bot">
                    <div class="found">
                        <el-button type="primary" class="cancel" @click="closeDia">{{$t('deal.Cancel')}}</el-button>
                        <el-button type="primary" @click="submitForm('ruleForm')">{{$t('deal.Submit')}}</el-button>
                    </div>
                    <!-- <a :href="found_link" target="_blank">{{$t('uploadFile.upload_funds')}}</a> -->
                </div>
            </div>
        </el-dialog>

        <el-dialog title="" :visible.sync="finishTransaction" :close-on-click-modal="false" :width="width"
            custom-class="completeDia">
            <img src="@/assets/images/alert-icon.png" class="resno" />
            <h1>{{$t('uploadFile.COMPLETED')}}!</h1>
            <h3>{{$t('uploadFile.SUCCESS')}}</h3>
            <a :href="baseAddressURL+'tx/'+txHash" target="_blank">{{txHash}}</a>
            <a class="a-close" @click="finishClose">{{$t('uploadFile.CLOSE')}}</a>
        </el-dialog>

        <el-dialog title="" :visible.sync="failTransaction" :width="width"
            custom-class="completeDia">
            <img src="@/assets/images/error.png" class="resno" />
            <h1>{{$t('uploadFile.Fail')}}!</h1>
            <h3>{{$t('uploadFile.FailTIP')}}</h3>
            <a :href="baseAddressURL+'tx/'+txHash" target="_blank">{{txHash}}</a>
            <a class="a-close" @click="failClose">{{$t('uploadFile.CLOSE')}}</a>
        </el-dialog>

        <el-dialog title="" :visible.sync="waitTransaction" :width="width"
            :before-close="failClose"
            custom-class="completeDia">
            <img src="@/assets/images/waiting.png" class="resno" />
            <h1>{{$t('uploadFile.waiting')}}</h1>
            <h3>{{$t('uploadFile.waitingTIP')}}</h3>
            <a :href="baseAddressURL+'tx/'+txHash" target="_blank">{{txHash}}</a>
            <a class="a-close" @click="failClose">{{$t('uploadFile.CLOSE')}}</a>
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
            <img src="@/assets/images/box-important.png" class="resno" />
            <h2>{{$t('uploadFile.file_uploaded')}}</h2>
            <h4>{{$t('uploadFile.file_uploaded_tip')}}</h4>
            <h4>{{$t('uploadFile.file_uploaded_tip01')}}</h4>
            <a class="a-close" @click="paymentPopup=false">{{$t('uploadFile.OK')}}</a>
        </el-dialog>

        <el-dialog title="" :visible.sync="paymentPopup01" :width="width"
            custom-class="completeDia">
            <img src="@/assets/images/box-important.png" class="resno" />
            <h2>{{$t('uploadFile.file_uploaded')}}</h2>
            <h4>{{$t('uploadFile.file_uploaded_tip02')}}</h4>
            <h4>{{$t('uploadFile.file_uploaded_tip01')}}</h4>
            <a class="a-close" @click="paymentPopup01=false">{{$t('uploadFile.OK')}}</a>
        </el-dialog>

        <el-dialog title="" :visible.sync="metamaskLoginTip" :width="width"
            custom-class="completeDia">
            <img src="@/assets/images/box-important.png" class="resno" />
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
                }else {
                    callback();
                }
            };
            return {
                tableData: {
                    file_name: '',
                    task_name: '',
                    payload_cid: ''
                },
                ruleForm: {
                    duration: '525',
                    storage_copy: '5',
                    fileList: [],
                    fileList_tip: false,
                    fileList_tip_text: '',
                    file_size: '',
                    file_size_byte: '',
                    storage_cost: '',
                    lock_plan: '2',
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
                biling_price: 0,
                storage: 0,
                storage_cost_low: 0,
                storage_cost_average: 0,
                storage_cost_high: 0,
                center_fail: false,
                centerDialogVisible: false,
                modelClose: true,
                width: document.body.clientWidth>600?'400px':'95%',
                widthUpload: document.body.clientWidth>600?'450px':'95%',
                widthDia: document.body.clientWidth<=600?'95%':document.body.clientWidth<=1440?'7rem':'6.6rem',
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
                waitTransaction: false,
                txHash: '',
                fileUploadVisible: false,
                paymentPopup: false,
                paymentPopup01: false,
                percentIn: '',
                loadMetamaskPay: false,
                metamaskLoginTip: false,
                lastTime: 0,
                found_link: process.env.NODE_ENV == "production"?"https://calibration-faucet.filswan.com/":"http://192.168.88.216:8080/faucet/#/dashboard",
                free: false
            };
        },
        props: ['uploadDigShow'],
        components: {},
        watch: {
            'ruleForm.fileList': function(){
                this.ruleForm.lock_plan = '2'
                this.$refs['ruleForm'].validate((valid) => {});
                if(this.ruleForm.file_size_byte > 0){
                    this.calculation()
                }
            }
        },
        methods: {
            calculation(type){
                this.ruleForm.storage_cost = this.ruleForm.file_size_byte * this.ruleForm.duration * this.storage / 365
                let _price = this.ruleForm.storage_cost * this.biling_price
                let number_price = Number(_price).toFixed(9)
                this.ruleForm.amount_minprice = number_price > 0.000000001 ? number_price : '0.0000000005'
                this.storage_cost_low = number_price > 0 ? Number(_price * 2).toFixed(9) : '0.000000001'
                this.storage_cost_average = number_price > 0 ? Number(_price * 3).toFixed(9) : '0.000000002'
                this.storage_cost_high = number_price > 0 ? Number(_price * 5).toFixed(9) : '0.000000003'
                this.ruleForm.amount = this.free ? 0 : this.storage_cost_average
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
            signFun(){
                let _this = this
                if(!_this.metaAddress || _this.metaAddress == 'undefined'){
                    NCWeb3.Init(addr=>{
                        _this.$nextTick(() => {
                            _this.$store.dispatch('setMetaAddress', addr)
                            _this.$emit("getMetamaskLogin", true)
                        })
                    })
                    return false
                }
            },
            closeDia() {
                this.$emit('getUploadDialog', false)
            },
            submitForm(formName) {
                let _this = this;
                _this.$refs[formName].validate((valid) => {
                    if (valid) {
                        if(_this.metaAddress&&!(_this.networkID==137)) {
                            _this.metamaskLoginTip = true
                            return false
                        }
                        if(!_this.ruleForm.lock_plan || (_this.ruleForm.amount<=0&&!_this.free)){
                            _this.ruleForm.lock_plan_tip = true
                            return false
                        }
                        if(_this.ruleForm.fileList_tip || isNaN(_this.ruleForm.amount)) return false

                        let now = new Date().valueOf(); //防止短时间内连续点击后多次触发请求的操作
                        if(_this.lastTime == 0){
                            _this.lastTime = now;
                        }else{
                            if((now - _this.lastTime) > 2000){
                                //重置上一次点击时间，设置的2秒间隔
                                _this.lastTime = now;
                            }else{
                                return false
                            }
                        }

                        if(_this.metaAddress){
                            // 授权代币
                            contract_erc20 = new web3.eth.Contract( erc20_contract_json );
                            contract_erc20.options.address = _this.usdcAddress
                            // 查询剩余代币余额为：
                            contract_erc20.methods.balanceOf(_this.metaAddress).call()
                            .then(resultUSDC => {
                                _this.usdcAvailable = web3.utils.fromWei(resultUSDC, 'ether');
                                console.log('Available:', _this.usdcAvailable, _this.ruleForm.amount)

                                // 判断支付金额是否大于代币余额
                                if(Number(_this.ruleForm.amount) > Number(_this.usdcAvailable) ){
                                    _this.$message.error('Insufficient balance')
                                    return false
                                }

                                // 通过 FormData 对象上传文件
                                var formData = new FormData()
                                formData.append('file', _this._file)
                                formData.append('duration', _this.ruleForm.duration)
                                formData.append('storage_copy', _this.ruleForm.storage_copy)
                                formData.append('wallet_address', _this.metaAddress)
                                // formData.append('task_name', _this.ruleForm.task_name)
                                _this.loading = true
                                _this.fileUploadVisible = true

                                let xhr = new XMLHttpRequest()
                                xhr.open("POST", `${_this.baseAPIURL}api/v1/storage/ipfs/upload`, true);   // 设置xhr得请求方式和url。
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
                                                _this.percentIn = xhr.readyState === 4 ? '100%' : _this.percentIn
                                                if(res.status == "success"){
                                                    if(res.data.status == 'Free') {
                                                        _this.finishClose()
                                                        return false
                                                    }else{
                                                        contract_erc20.methods.allowance(_this.gatewayContractAddress, _this.metaAddress).call()
                                                        .then(resultUSDC => {
                                                            console.log('allowance：'+ resultUSDC);
                                                            if(resultUSDC < web3.utils.toWei(_this.ruleForm.amount, 'ether')){
                                                                contract_erc20.methods.approve(_this.gatewayContractAddress, web3.utils.toWei(_this.ruleForm.amount, 'ether')).send({from:  _this.metaAddress})
                                                                .then(receipt => {
                                                                    // console.log('approve receipt:', receipt)
                                                                    _this.contractSend(res.data)
                                                                })
                                                                .catch(error => {
                                                                    // console.log('errorerrorerror', error)
                                                                    _this.finishClose()
                                                                })
                                                            }else{
                                                                _this.contractSend(res.data)
                                                            }
                                                        })
                                                    }
                                                }else{
                                                    _this.$message.error(_this.$t('uploadFile.xhr_tip'))
                                                }
                                            }
                                        }
                                    } else {
                                        i += 1
                                        if(i <= 1){
                                            if (xhr.status === 500) _this.$message.error(_this.$t('uploadFile.xhr_error500'))
                                        }
                                        _this.loading = false
                                        _this.fileUploadVisible = false
                                    }
                                    xhr.upload.addEventListener("error", event => {
                                        _this.$message.error(_this.$t('uploadFile.xhr_tip'))
                                    })

                                    xhr.upload.addEventListener("progress", event => {
                                        if (event.lengthComputable) {
                                            let loaded = event.loaded
                                            let total = event.total
                                            // console.log('total-loaded', total, loaded)
                                            let percentIn = Math.floor(event.loaded / event.total * 100);
                                            // 设置进度显示
                                            _this.percentIn = percentIn+'%'
                                            // console.log(percentIn+'%-')
                                        }
                                    })
                                };
                                // 获取上传进度
                                xhr.upload.onprogress = function(event) { 
                                    // console.log('event.loaded', event.loaded)
                                    // console.log('event.total', event.total)
                                    if (event.lengthComputable) {
                                        let percentIn = Math.floor(event.loaded / event.total * 100);
                                        // 设置进度显示
                                        if(percentIn == 100 && event.loaded <= event.total) percentIn = 99
                                        _this.percentIn = percentIn == 100 && event.loaded <= event.total ? '99%' : percentIn+'%'
                                        // console.log(percentIn+'%')
                                    }
                                };
                                xhr.send(formData);
                                return false
                            })
                        }
                    } else {
                        console.log('error submit!!');
                        return false;
                    }
                });
            },
            contractSend(resData){
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
                    id: resData.w_cid,
                    minPayment: web3.utils.toWei(String(_this.ruleForm.amount_minprice), 'ether'),
                    amount: web3.utils.toWei(_this.ruleForm.amount, 'ether'),
                    lockTime: 86400 * Number(_this.$root.LOCK_TIME), // one day
                    recipient: _this.recipientAddress, //todo:
                    size: resData.file_size,
                    copyLimit: Number(_this.ruleForm.storage_copy)
                }
                console.log(lockObj)
                contract_instance.methods.lockTokenPayment(lockObj)
                .send(payObject)
                .on('transactionHash', function(hash){
                    // console.log('hash console:', hash);
                    _this.loadMetamaskPay = true
                    _this.txHash = hash
                })
                .on('confirmation', function(confirmationNumber, receipt){
                    // console.log('confirmationNumber console:', confirmationNumber, receipt);
                })
                .on('receipt', function(receipt){
                    // receipt example
                    console.log('receipt console:', receipt);
                    _this.checkTransaction(receipt.transactionHash, resData, lockObj)
                    _this.txHash = receipt.transactionHash
                })
                .on('error', function(error){
                    console.log('lockTokenPayment error console:', error)
                    // console.error
                    _this.loading = false
                    _this.loadMetamaskPay = false
                    if(_this.finishTransaction) return false
                    if(_this.txHash) _this.waitTransaction = true
                    else _this.failTransaction = true
                }); 
            },
            failClose(){
                this.failTransaction = false
                this.waitTransaction = false
                this.txHash = ''
            },
            checkTransaction(txHash, resData, lockObj) {
                let _this = this
                web3.eth.getTransactionReceipt(txHash).then(
                    async res => {
                        console.log('checking ... ');
                        if (!res) { return _this.timer = setTimeout(() => { _this.checkTransaction(txHash, resData, lockObj); }, 2000); }
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
                this.txHash = ''
                this.$emit('getUploadDialog', false, true)
            },
            // 文件上传
            uploadFile(params) {
                this._file = params.file;
                const isLt2M = this._file.size / 1024 / 1024 / 1024 <= 25;  // or 1000
                this.ruleForm.file_size = this.sizeChange(this._file.size)
                this.ruleForm.file_size_byte = this.byteChange(this._file.size)
                this.free = (this.free_quota_per_month - this.free_usage) >= this._file.size?true:false
                // console.log('bytes', this._file.size, this.free_quota_per_month-this.free_usage, this.free)
                if (!isLt2M) {
                    // this.$message.error(this.$t('deal.upload_form_file_tip'))
                    this.ruleForm.fileList_tip = true
                    this.ruleForm.fileList_tip_text = this.$t('deal.upload_form_file_tip')
                    return false
                }else{
                    this.ruleForm.fileList_tip = false
                }
            },
            sizeChange(bytes){
                if (bytes === 0) return '0 B';
                if (!bytes) return "-";
                var k = 1024, // or 1000
                    sizes = ['bytes', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB'],
                    i = Math.floor(Math.log(bytes) / Math.log(k));

                if (Math.round((bytes / Math.pow(k, i))).toString().length > 3) {
                    // 判断大小是999999999左右，解决会显示成1.00e+3科学计数法
                    i += 1
                }

                // if(i == 2) return (bytes / Math.pow(k, i)).toPrecision(3) + ' ' + sizes[i];
                return Number(bytes / Math.pow(k, i)) + ' ' + sizes[i];
            },
            byteChange(limit){
                var size = "";
                // 只转换成GB
                if(limit <= 0){
                    return '-'
                }else{
                    size = limit/( 1024 * 1024 * 1024)  //or 1000
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
            async stats(){
                let _this = this
                _this.loading = true
                if(_this.$root.SWAN_PAYMENT_CONTRACT_ADDRESS){
                    _this.gatewayContractAddress = _this.$root.SWAN_PAYMENT_CONTRACT_ADDRESS
                    _this.usdcAddress = _this.$root.USDC_ADDRESS
                    _this.recipientAddress = _this.$root.RECIPIENT
            
                    const storageRes = await _this.sendRequest(`${process.env.BASE_API}stats/storage?wallet_address=${_this.metaAddress}`)
                    let cost = storageRes.data.average_price_per_GB_per_year?storageRes.data.average_price_per_GB_per_year.split(" "):[]
                    if(cost[0]) _this.storage = cost[0]

                    _this.biling_price = _this.$root.filecoin_price

                    _this.loading = false
                }else {
                    setTimeout(function(){
                        _this.stats()
                    }, 1000)
                }
            },
            async sendRequest(apilink) {
                try {
                    const response = await axios.get(apilink)
                    return response.data
                } catch (err) {
                    console.error(err)
                }
            }
        },
        mounted() {
            let _this = this
            that = _this
            _this.stats()
            // _this.$store.dispatch("setRouterMenu", 0);
            // _this.$store.dispatch('setHeadertitle', _this.$t('navbar.Upload_files'))
            // document.onkeydown = function(e) {
            //     if (e.keyCode === 13) {
            //     }
            // }
        },
        filters: {
            NumStorage(value) {
                if (String(value) === '0') return 0;
                if (!value) return "-";
                return value.toFixed(10);
            },
            NumStoragePlan(value) {
                if (!value) return "-";
                return value.toFixed(9) > 0 ? value.toFixed(9) : '0.000000001';
            },
            byteStorage(limit) {
                // 只转换成GB
                if(limit <= 0){
                    return '0'
                }else{
                    return limit/( 1024 * 1024 * 1024)  //or 1000
                }
            }
        },
        computed: {
            metaAddress() {
                return this.$store.getters.metaAddress
            },
            networkID() {
                return this.$store.getters.networkID
            },
            free_usage() {
                return this.$store.getters.free_usage
            },
            free_quota_per_month() {
                return this.$store.getters.free_quota_per_month
            }
        },
    };
</script>


<style scoped lang="scss">
    .el-dialog__wrapper /deep/ {  
        display: flex;
        align-items: center;
        .metaM{
            margin: auto !important;
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
                    font-size: 0.18rem;
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
                    font-size: 0.18rem;
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
            margin: auto !important;
            text-align: center;
            box-shadow: 0 0 13px rgba(128,128,128,0.8);
            border-radius: 0.2rem;
            .el-dialog__header{
            display: none;
            }
            img{
                display: block;
                max-width: 100px;
                margin: 0 auto 0.3rem;
            }
            h1{
                margin: 0.22rem auto 0.1rem;
                font-size: 0.32rem;
                font-weight: 500;
                line-height: 1.2;
                color: #191919;
                word-break: break-word;
            }
            h2{
                margin: 0 auto 0.1rem;
                font-size: 0.22rem;
                font-weight: 600;
                line-height: 1.4;
                color: #555;
                word-break: break-word;
                text-align: center;
            }
            h4{
                font-size: 0.2rem;
                font-weight: 500;
                line-height: 1.4;
                color: #555;
                word-break: break-word;
                text-align: center;
            }
            h3, a{
                font-size: 0.22rem;
                font-weight: 500;
                line-height: 1.4;
                color: #191919;
                word-break: break-word;
            }
            a{
                text-decoration: underline;
                color: #007bff;
            }
            a.a-close{
                height: 0.6rem;
                line-height: 0.6rem;
                padding: 0 45px;
                background: linear-gradient(45deg,#4f8aff, #4b5eff);
                font-size: 0.22rem;
                color: #fff;
                border-radius: 0.14rem;
                cursor: pointer;
                margin: 0.4rem auto 0;
                display: block;
                width: max-content;
                text-decoration: unset;
            }
        }
        .fileUpload{
            margin: auto !important;
            box-shadow: 0 0 13px rgba(128,128,128,0.8);
            border-radius: 0.2rem;
            .el-dialog__header{
                padding: 0.3rem 0.4rem;
                color: #000;
                font-size: 0.22rem;
                font-weight: 500;
                line-height: 1;
                text-transform: capitalize;
            }
            .el-dialog__body{
                padding: 0 0.4rem;
                h3{   
                    margin: 0 0 0.1rem;
                    font-size: 0.2rem;
                    font-weight: normal;
                    line-height: 1.4;
                    color: #555;
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
        .uploadDig{
            background: #fff;
            margin: auto !important;
            box-shadow: 0 0 13px rgba(128,128,128,0.8);
            border-radius: 0.2rem;
            .el-dialog__header{
                padding: 0.3rem 0.4rem;
                display: flex;
                align-items: center;
                border-bottom: 1px solid #dfdfdf;
                color: #000;
                font-size: 0.22rem;
                font-weight: 500;
                line-height: 1;
                text-transform: capitalize;
                @media screen and (max-width: 479px){
                    padding: 0.3rem 0.2rem;
                }
                img{
                    width: 20px;
                    height: 20px;
                    margin: 0 0 0 5px;
                    cursor: pointer;
                    @media screen and (max-width:1440px){
                        width: 17px;
                        height: 17px;
                    }
                    @media screen and (max-width: 1280px){
                        width: 16px;
                        height: 16px;
                    }
                }
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
                padding: 0 0.4rem;
                @media screen and (max-width: 479px){
                    padding: 0 0.2rem;
                }
                .upload_form{
                    // display: flex;
                    // align-items: baseline;
                    width: 100%; 
                    margin: auto; 
                    justify-content: flex-start;
                    .el-form{
                        width: 100%;
                        margin: 0;
                        .el-form-item::after, .el-form-item::before{
                            display: none;
                        }
                        .el-form-item{
                            display: flex;
                            align-items: center;
                            justify-content: space-between;
                            width: 100%;
                            margin: 0.2rem auto;
                            .el-form-item__label{
                                display: flex;
                                justify-content: flex-start;
                                align-items: center;
                                width: 47%;
                                padding: 0 2% 0 0;
                                // max-width: 2rem;
                                line-height: 1.5;
                                text-align: left;
                                font-size: 0.2rem;
                                white-space: normal;
                                color: #000;
                                font-weight: 500;
                                text-shadow: 0 0 black;
                                text-align: right;
                                @media screen and (max-width: 479px){
                                    padding: 0;
                                }
                                img{
                                    width: 20px;
                                    height: 20px;
                                    margin: 0 0 0 5px;
                                    cursor: pointer;
                                    @media screen and (max-width:1440px){
                                        width: 17px;
                                        height: 17px;
                                    }
                                    @media screen and (max-width: 1280px){
                                        width: 16px;
                                        height: 16px;
                                    }
                                }
                                &::before{
                                    display: none;
                                }
                            }
                            .el-form-item__content{
                                width: 50%;
                                display: flex;
                                align-items: center;
                                justify-content: flex-end;
                                font-size: 0.2rem;
                                white-space: normal;
                                word-break: break-word;
                                line-height: 1.5;
                                color: #555;
                                h4{
                                    width: 100%;
                                    font-size: 0.1372rem;
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
                                }
                                .el-tag, .el-button--small{
                                    margin: 0 5px 5px 0;
                                }
                                .el-input{
                                    width: auto;
                                    .el-input__inner{
                                        height: 0.32rem;
                                        font-size: 0.1372rem;
                                        line-height: 0.32rem;
                                    }
                                    .el-input__suffix{
                                        display: none;
                                    }
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
                                            // height: 0.32rem;
                                            // padding: 0 0.2rem;
                                            // margin: 0 5px 0 0;
                                            // line-height: 0.32rem;
                                            // background-color:transparent;
                                            // border: 1px solid #2c4c9e;
                                            // border-radius: 0.08rem;
                                            // color: #2c4c9e;
                                            // font-size: 0.1372rem;
                                            display: flex;
                                            align-items: center;
                                            justify-content: center;
                                            height: 0.5rem;
                                            padding: 0 0.2rem;
                                            margin: 0;
                                            background: linear-gradient(45deg,#4e88ff, #4b5fff);
                                            border-radius: 0.14rem;
                                            line-height: 0.5rem;
                                            text-align: center;
                                            color: #fff;
                                            font-size: 0.18rem;
                                            font-family: inherit;
                                            border: 0;
                                            outline: none;
                                            transition: background-color .3s, border-color .3s, color .3s, box-shadow .3s;
                                            cursor: pointer;
                                            span{
                                                display: flex;
                                                align-items: center;
                                            }
                                            img{
                                                display: inline-block;
                                                height: 20px;
                                                margin: 0 0.1rem 0 0;
                                                @media screen and (max-width: 1280px){
                                                    height: 16px;
                                                }
                                            }
                                            &:hover{
                                                opacity: .9;
                                                box-shadow: 0 12px 12px -12px rgba(12, 22, 44, 0.32);
                                            }
                                        }
                                    }
                                    .el-upload-list{
                                        width: 100%;
                                        max-width: 300px;
                                        float: none;
                                        clear: both;
                                        @media screen and (max-width: 1440px){
                                            max-width: 250px;
                                        }
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
                    width: 100%;
                    margin: auto; 
                    justify-content: flex-start;
                    .title{
                        display: flex;
                        align-items: center;
                        justify-content: flex-start;
                        margin: 0;
                        line-height: 1.5;
                        text-align: center;
                        font-size: 0.22rem;
                        white-space: normal;
                        color: #000;
                        font-weight: 500;
                        // text-shadow: 0 0 black;
                        text-indent: 0;
                        img{
                            width: 20px;
                            height: 20px;
                            margin: 0 0 0 5px;
                            cursor: pointer;
                            @media screen and (max-width:1440px){
                                width: 17px;
                                height: 17px;
                            }
                            @media screen and (max-width: 1280px){
                                width: 16px;
                                height: 16px;
                            }
                        }
                    }
                    .desc{
                        margin: 0 0 0.1rem;
                        line-height: 1.5;
                        font-size: 0.16rem;
                        white-space: normal;
                        color: #999;
                        font-weight: normal;
                    }
                    .upload_plan_radio{
                        .el-radio-group{
                            width: 100%;
                            background: #f7f7f7;
                            border-radius: 0.2rem;
                            .el-radio{
                                display: flex;
                                align-items: center;
                                justify-content: space-between;
                                width: 100%;
                                height: auto;
                                padding: 0.2rem 0.3rem;
                                margin: auto;
                                border: 0;
                                // line-height:30px;
                                .el-radio__input{    
                                    width: 20px;
                                    display: flex;
                                    align-items: center;
                                    .el-radio__inner{
                                        border-color: #555;
                                    }
                                }
                                .el-radio__input.is-checked{
                                    .el-radio__inner{
                                        position: relative;
                                        width: 16px;
                                        height: 16px;
                                        border-color: transparent;
                                        background: transparent;
                                        &:after {
                                            content: "";
                                            display: block;
                                            height: 16px;
                                            width: 16px;
                                            background-image: url(../assets/images/icon_xuanzhong@2x.png);
                                            background-size: 100%;
                                            position: absolute;
                                            left:0;
                                            top:0;
                                            transform: translate(0, 0) scale(1);
                                            transition: all 0.15s;
                                        }
                                    }
                                }
                                .el-radio__label{
                                    display: flex;
                                    justify-content: space-between;
                                    width: calc(100% - 30px);
                                    .title{
                                        font-size: 0.2rem;
                                        line-height: 1;
                                    }
                                    .cont{
                                        font-size: 0.2rem;
                                        font-weight: 500;
                                        line-height: 1;
                                        text-align: center;
                                    }
                                }
                            }
                            .el-radio:nth-child(3n+1){
                                .el-radio__label{
                                    .cont{
                                        color: #35AD92;
                                    }
                                }
                            }
                            .el-radio:nth-child(3n+2){
                                border-top: 1px solid #dfdfdf;
                                border-bottom: 1px solid #dfdfdf;
                                .el-radio__label{
                                    .cont{
                                        color: #2C7FF8;
                                    }
                                }
                            }
                            .el-radio:nth-child(3n+3){
                                .el-radio__label{
                                    .cont{
                                        color: #F63D3D;
                                    }
                                }
                            }
                            .el-radio:hover{
                                background-color: rgba(64,158,255,0.1);
                            }

                        }

                    }
                    .upload_plan_radio_free{
                        opacity: .5;
                    }
                }
                .upload_bot{
                    width: 100%;
                    margin: 0.25rem auto 0.2rem;
                    text-align: center;
                    .found{
                        display: flex;
                        justify-content: space-between;
                        align-items: center;
                        width: 100%;
                        text-align: center;
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
                    a{
                        margin: auto;
                        text-decoration: underline;
                        font-size: 0.18rem;
                        color: rgb(11, 49, 143);
                        cursor: pointer;
                    }
                }
            }
            .dialog-footer{
                display: flex;
                align-items: center;
                justify-content: flex-end;
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
            border-radius: 0.2rem;
            .el-loading-spinner{
                top: 0;
                position: relative;
                margin: 0 0 0.2rem;
            }
            p{
                font-size: 14px;
                color: #555;
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
                font-size: 0.2rem;
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
                        .el-form-item__label{
                            display: flex;
                            justify-content: flex-end;
                            align-items: center;
                            width: 47%;
                            padding: 0 2% 0 0;
                            // max-width: 2rem;
                            line-height: 1.5;
                            text-align: left;
                            font-size: 0.18rem;
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
                        }
                        .el-form-item__content{
                            display: flex;
                            align-items: center;
                            font-size: 0.18rem;
                            white-space: normal;
                            word-break: break-word;
                            line-height: 1.5;
                            color: #666;
                            @media screen and (max-width:600px){
                                font-size: 14px;
                            }
                            small{
                                margin: 2px 5px 0;
                                font-size: 12px;
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
                                    height: 0.32rem;
                                    font-size: inherit;
                                    line-height: 0.32rem;
                                }
                                .el-input__suffix{
                                    display: none;
                                }
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
                                        font-size: 0.18rem;
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
                    font-size: 0.2rem;
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
                    font-size: 0.18rem;
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
                                    font-size: 0.18rem;
                                }
                                .cont{
                                    font-size: 0.18rem;
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
                .upload_plan_radio_free{
                    opacity: .5;
                }
            }
            .upload_bot{
                display: flex;
                justify-content: center;
                align-items: center;
                flex-wrap: wrap;
                width: 100%;
                margin: 0.25rem auto 0.15rem;
                .el-button /deep/{
                    height: 0.45rem;
                    padding: 0 0.4rem;
                    margin-left: 0;
                    background-color: #0b318f;
                    line-height: 0.45rem;
                    font-size: 0.18rem;
                    color: #fff;
                    border: 0;
                }
                .no_login{
                    background: #f56c6c;
                }
                .found{
                    width: 100%;
                    text-align: center;
                    a{
                        text-decoration: underline;
                        font-size: 0.18rem;
                        color: rgb(11, 49, 143);
                        cursor: pointer;
                    }
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
                        font-size: 0.18rem;
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
    @media screen and (max-width: 1024px) {
        #Create{
            .upload {
                .upload_form {
                    .el-form {
                        width: 48%;
                    }
                }
            }
        } 
    }
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
                     .el-form /deep/{
                         .el-form-item{
                             flex-wrap: wrap;
                             .el-form-item__label{
                                width: 100%;
                                padding-bottom: 0.05rem;
                                justify-content: center;
                             }
                             .el-form-item__content{
                                flex-wrap: wrap;
                                width: 100%;
                                justify-content: center;
                             }
                         }
                     }
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
