<template>
        <el-dialog :title="$t('billing.download_module_title')+titlePage" :modal="true" :width="widthDia" :visible.sync="downVisible" :before-close="closeDia">
            <div class="upload_form">
                <el-date-picker
                    v-model="downloadTime"
                    type="daterange"
                    :range-separator="$t('billing.time_to')"
                    :start-placeholder="$t('billing.start_date')"
                    :end-placeholder="$t('billing.end_date')">
                </el-date-picker>
                <div class="drag_container">
                    <span>{{$t('billing.verify')}}</span>
                    <drag-verify 
                        ref="Verify"
                        :width="width" 
                        :height="height" 
                        :text="text" 
                        :success-text="successText" 
                        :background="background" 
                        :handlerBg="handlerBg" 
                        :progress-bar-bg="background" 
                        :completed-bg="completedBg" 
                        :handler-icon="handlerIcon" 
                        :text-size="textSize" 
                        :success-icon="successIcon"
                        :circle="false" >
                    </drag-verify>
                </div>
                <div class="upload_bot">
                    <el-button type="primary" @click="download">{{$t('billing.download_module_btn')}}</el-button>
                </div>
            </div>
        </el-dialog>
</template>

<script>
import dragVerify from "vue-drag-verify";
export default {
  name: "download",
  data() {
    return {
      widthDia: document.body.clientWidth <= 600 ? "95%" : document.body.clientWidth <= 1600 ? "520px" : "640px",
      downloadTime: "",

      handlerIcon: "el-icon-d-arrow-right",
      successIcon: "el-icon-check",
      background: "#dadada",
      handlerBg: "#fff",
      completedBg: "#66cc66",
      text: this.$t("billing.verify_tip"),
      successText: this.$t("billing.verify_result"),
      width:
        document.body.clientWidth > 1600
          ? 480
          : document.body.clientWidth > 600 ? 400 : 290,
      height:
        document.body.clientWidth > 1800
          ? 55
          : document.body.clientWidth > 1440
            ? 48
            : document.body.clientWidth > 1280 ? 40 : 35,
      textSize: "0.22rem",
      textColor: "#888",
      isCircle: "true",
      radius: "12px"
    };
  },
  props: ["downVisible", 'titlePage'],
  components: {
    dragVerify
  },
  methods: {
    closeDia() {
      this.$emit("getDownload", false);
    },
    download() {
      // console.log(this.$refs.Verify.isMoving)
      console.log(
        "Time:",
        this.downloadTime,
        "Verify:",
        this.$refs.Verify.isPassing
      );
      // console.log(this.$refs.Verify.handlerIconClass)
    }
  },
  mounted() {}
};
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
      padding: 0.3rem 0.4rem;
      display: flex;
      border-bottom: 1px solid #dfdfdf;
      .el-dialog__title {
        color: #333;
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
      padding: 0.3rem 0.4rem 0.1rem;
      .el-range-editor.el-input__inner {
        position: relative;
        width: 100%;
        height: auto;
        padding: 0;
        line-height: 0.6rem;
        border-color: #888;
        font-size: 0.22rem;
        color: #333;
        flex-wrap: wrap;
        border: 0;
        justify-content: center;
        i {
          display: none;
        }
        .el-icon-date {
          display: block;
          position: absolute;
          top: 0;
          left: 0;
          width: 0.8rem;
          height: 0.62rem;
          border-right: 1px solid #888;
          margin: 0;
        }
        .el-range-input {
          width: 100%;
          height: 0.6rem;
          font-size: inherit;
          font-family: inherit;
          border: 1px solid #888;
          border-radius: 0.14rem;
          &:before {
            // content:"\e78e";
            content: "`";
          }
        }
        .el-range__icon,
        .el-range-separator {
          line-height: inherit;
          display: flex;
          align-items: center;
          justify-content: center;
          font-size: inherit;
          color: #000;
        }
      }
      .drag_container {
        display: flex;
        align-items: center;
        justify-content: space-between;
        margin: 0.2rem 0 0;
        font-size: 0.22rem;
        color: #888;
        .drag_verify {
          border-radius: 0.12rem !important;
          .dv_progress_bar {
            width: 100% !important;
            border-radius: 0.12rem !important;
          }
          .dv_text {
            color: #888;
          }
          .dv_handler {
            display: flex;
            align-items: center;
            justify-content: center;
            top: 4px;
            bottom: 4px;
            height: auto !important;
            margin-left: 4px;
            border-radius: 0.12rem !important;
            border: 0;
            i {
              color: #dadada;
              font-size: 0.23rem;
            }
          }
        }
      }
      .upload_bot {
        display: flex;
        justify-content: space-between;
        align-items: center;
        width: 100%;
        margin: 0.25rem auto 0.2rem;
        .el-button {
          width: 100%;
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
        }
      }
    }
    .dialog-footer {
      display: flex;
      align-items: center;
      justify-content: flex-end;
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
