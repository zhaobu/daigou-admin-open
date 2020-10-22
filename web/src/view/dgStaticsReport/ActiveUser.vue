<template>
  <div class="chartContainer">
    <el-form
      ref="elFormRef"
      :model="formData"
      size="medium"
      label-position="left"
      label-width="70px"
      :inline="true"
      :rules="formRules"
    >
      <!-- 选择统计类型 -->
      <el-form-item label="统计类型" class="statType">
        <el-select
          v-model="formData.statTypeValue"
          size="medium"
          placeholder="请选择"
          @change="onStatSelectChange"
          :style="{width: '70px'}"
        >
          <el-option
            v-for="item in formData.statTypeOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          ></el-option>
        </el-select>
      </el-form-item>
      <!-- 选择时间段 -->
      <el-form-item label="时间段" class="timePicker" required>
        <el-col :span="8">
          <el-form-item prop="timeStartValue">
            <el-date-picker
              v-model="formData.timeStartValue"
              :type="formData.timePickerType"
              align="center"
              placeholder="开始时间"
              :style="{width: '135px'}"
              :clearable="false"
            ></el-date-picker>
          </el-form-item>
        </el-col>
        <el-col class="datePickerLine" :span="2">至</el-col>
        <el-col :span="8">
          <el-form-item prop="timeEndValue">
            <el-date-picker
              v-model="formData.timeEndValue"
              :type="formData.timePickerType"
              align="center"
              placeholder="结束时间"
              :style="{width: '135px'}"
              :clearable="false"
            ></el-date-picker>
          </el-form-item>
        </el-col>
      </el-form-item>
      <!-- 查询按钮 -->
      <el-form-item>
        <el-button type="primary" :disabled="datePickerReady" @click="onGetData('elFormRef')">查询</el-button>
      </el-form-item>
    </el-form>
    <v-chart :theme="theme" :autoresize="autoresize" :options="showData" />
  </div>
</template>

<script>
import ECharts from "vue-echarts";
// echarts-gl包含大部分图表所需要的依赖模块
import "echarts-gl";
import { macarons } from "./themes/macarons.js";
// import { ovilia_green } from "./themes/ovilia_green.js";
import { compareDate, timeToGolang, formatTimeToStr } from "@/utils/date";
import { activeUser } from "@/api/daigou/staticsReport";

// 按需引入不同的图表需要的依赖
// import "echarts/lib/chart/line";
// import "echarts/lib/component/polar";
// import "echarts/lib/chart/gauge";
// registering custom theme
const stat_type = { day: 0, week: 1, month: 2 };
const legendName = ["日活跃用户数", "周活跃用户数", "月活跃用户数"];
export default {
  name: "ActiveUser",
  components: {
    "v-chart": ECharts,
  },
  computed: {
    datePickerReady() {
      if (
        this.formData.timeStartValue === "" ||
        this.formData.timeEndValue === ""
      ) {
        return true;
      } else if (
        compareDate(this.formData.timeStartValue, this.formData.timeEndValue)
      ) {
        return true;
      }
      return false;
    },
  },
  data() {
    const validStart = (rule, value, callback) => {
      // console.error(this.formData.timeStartValue);
      // console.error({ name: "validStart", rule, value, callback });
      if (value === "") {
        callback(new Error("请输入开始时间"));
      } else {
        if (this.formData.timeEndValue !== "") {
          this.$refs.elFormRef.validateField("timeEndValue");
        }
        callback();
      }
    };

    const validEnd = (rule, value, callback) => {
      // console.error({ name: "validEnd", rule, value, callback });
      if (value === "") {
        callback(new Error("请输入结束时间"));
      } else if (this.datePickerReady) {
        callback(new Error("结束时间需晚于开始时间"));
      } else {
        callback();
      }
    };

    const lastMonth = () => {
      const newTime = new Date();
      newTime.setTime(newTime.getTime() - 3600 * 1000 * 24 * 30);
      return newTime;
    };
    return {
      //在组件根元素尺寸变化时是否需要自动进行重绘
      autoresize: true,
      //主题
      theme: macarons,

      formData: {
        //当前选择的统计类型
        statTypeValue: 0,
        //可供选择统计类型
        statTypeOptions: [
          {
            value: stat_type.day,
            label: "天",
          },
          {
            value: stat_type.week,
            label: "周",
          },
          {
            value: stat_type.month,
            label: "月",
          },
        ],
        //开始时间
        timeStartValue: lastMonth(),
        //结束时间
        timeEndValue: new Date(),
        //时间选择器类型
        timePickerType: "date",
      },

      //校验规则
      formRules: {
        timeStartValue: [{ validator: validStart, trigger: "blur" }],
        timeEndValue: [{ validator: validEnd, trigger: "blur" }],
      },

      //默认展示数据
      showData: {
        title: {
          text: "活跃用户",
          align: "right",
        },
        tooltip: {
          trigger: "axis",
        },
        legend: {
          data: ["日活跃用户数"],
        },
        grid: {
          left: "3%",
          right: "4%",
          bottom: "3%",
          containLabel: true,
        },
        toolbox: {
          feature: {
            saveAsImage: {},
          },
        },
        xAxis: {
          type: "category",
          boundaryGap: false,
          data: ["周一", "周二", "周三", "周四", "周五", "周六", "周日"],
        },
        yAxis: {
          name: "人数",
          type: "value",
        },
        series: [
          {
            name: "日活跃用户数",
            type: "line",
            stack: "总量",
            data: [120, 132, 101, 134, 90, 230, 210],
            label: {
              normal: {
                show: true,
                position: "top",
              },
            },
          },
        ],
      },
    };
  },
  methods: {
    onStatSelectChange() {
      this.formData.timeStartValue = "";
      this.formData.timeEndValue = "";
      this.$refs.elFormRef.resetFields();
      switch (this.formData.statTypeValue) {
        case stat_type.week:
          this.formData.timePickerType = "week";
          break;
        case stat_type.month:
          this.formData.timePickerType = "month";
          break;
        default:
          this.formData.timePickerType = "date";
          break;
      }
    },

    async onGetData(formName) {
      this.$refs[formName].validate(async (valid) => {
        if (!valid) {
          this.$message.error(`请检查参数`);
          return;
        }

        let start_time = timeToGolang(this.formData.timeStartValue);
        let end_time = timeToGolang(this.formData.timeEndValue);
        //发生api请求
        const { code, data, msg } = await activeUser({
          query_type: this.formData.statTypeValue,
          start_time,
          end_time,
        });
        if (code != 0) {
          console.error({ code, data, msg });
          return;
        }
        this.setShowData(data);
      });
    },
    //绘制表格数据
    setShowData(data) {
      const index = 0;
      const xAxis = this.showData.xAxis;
      const series = this.showData.series[index];
      xAxis.data = [];
      series.data = [];
      for (const v of data) {
        if (this.formData.statTypeValue == stat_type.day) {
          xAxis.data.push(formatTimeToStr(v.register_time, "yyyy年-MM月-dd日"));
        } else if (this.formData.statTypeValue == stat_type.week) {
          xAxis.data.push(v.register_time + "周");
        } else if (this.formData.statTypeValue == stat_type.month) {
          xAxis.data.push(v.register_time + "月");
        }
        series.data.push(v.num);
      }
      //修改显示名称
      this.showData.legend.data[index] =
        legendName[this.formData.statTypeValue];
      this.showData.series[index].name =
        legendName[this.formData.statTypeValue];
    },
  },
  mounted() {
    this.onGetData("elFormRef");
  },
  watch: {},
};
</script>

<style scoped lang="scss">
.chartContainer {
  position: relative;
  width: 100%;
  height: calc(100vh - 84px);
}
.echarts {
  position: relative;
  margin-left: -20px;
  width: 98%;
  height: 88%;
}

.timePicker {
  margin-left: 50px;
  width: 420px;
  .datePickerLine {
    margin-left: 35px;
  }
}

.statType {
  margin-left: -13px;
}
</style>
