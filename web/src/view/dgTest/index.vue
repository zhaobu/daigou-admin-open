<template>
  <div class="app-container">
    <el-form ref="elFormRef" label-position="left" :inline="true">
      <el-form-item label="连接选项">
        <el-select ref="connectRef" v-model="connectVal">
          <el-option
            v-for="item in connectOptions"
            :key="item.value"
            :label="connectLabel(item.value)"
            :value="item.value"
          />
        </el-select>
        <el-form-item :style="{ 'margin-left': '20px'}">
          <el-button type="primary" size="medium" @click="hadelConnectChange">切换</el-button>
        </el-form-item>
      </el-form-item>
    </el-form>

    <el-divider></el-divider>

    <el-tabs>
      <el-tab-pane label="数据库">
        <mysql-connect :formData="formMysqlData"></mysql-connect>
      </el-tab-pane>
      <el-tab-pane label="Redis">
        <redis-connect :formData="formRedisData"></redis-connect>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>
<script>
import { getConnectList, changeConnect } from "@/api/daigou/serverConfig";
import Redis from "./components/Redis";
import Mysql from "./components/Mysql";

export default {
  name: "Mysql",
  components: {
    "mysql-connect": Mysql,
    "redis-connect": Redis,
  },
  props: [],
  data() {
    return {
      // redis配置信息
      redisList: null,
      // mysql配置
      mysqlList: null,
      //当前选中的连接
      connectVal: undefined,

      // 连接选项
      connectOptions: [],

      //当前连接的索引
      currentIndex: undefined,
    };
  },
  computed: {
    connectLabel() {
      return function (index) {
        let label = this.connectOptions[index].label;
        if (index == this.currentIndex) {
          label += "(当前连接)";
        }
        return label;
      };
    },
    formMysqlData() {
      if (!this.mysqlList) {
        return {};
      }
      const info = this.mysqlList[this.connectVal];

      //解析url得到ip和port
      let addr_info = info.addr.split(":");
      const { name, password, db_name, user_name, params } = info;
      const conf = {
        name,
        password,
        db_name,
        user_name,
        params: JSON.stringify(params),
        ip: addr_info[0],
        port: addr_info[1],
      };
      return conf;
    },
    formRedisData() {
      if (!this.redisList) {
        return {};
      }
      const info = this.redisList[this.connectVal];

      //解析url得到ip和port
      let addr_info = info.addr.split(":");
      const { name, password, db } = info;
      const conf = { name, password, db, ip: addr_info[0], port: addr_info[1] };

      return conf;
    },
  },
  created() {
    this.fetchData();
  },
  methods: {
    async fetchData() {
      const { code, data, msg } = await getConnectList();
      if (code != 0) {
        this.$message.error(`配置列表获取错误,${msg}`);
        return;
      }

      //处理连接选项
      this.connectOptions = [];
      const mysqlList = data.mysql_config_list;
      for (const index in mysqlList) {
        let label = mysqlList[index].name;
        let value = mysqlList[index].index;
        this.connectOptions.push({ label, value });
      }

      this.connectVal = data.current_index;
      this.currentIndex = data.current_index;
      this.mysqlList = data.mysql_config_list;
      this.redisList = data.redis_config_list;
    },

    async hadelConnectChange() {
      if (this.connectVal == this.currentIndex) {
        this.$message.warning(`要切换的连接未修改`);
        return;
      }

      const index = this.connectVal;
      //请求服务器切换
      const { code, msg } = await changeConnect({ index });
      if (code != 0) {
        this.$message.error(`配置列表获取错误,${msg}`);
        return;
      }

      this.fetchData();
    },
  },
};
</script>
<style>
</style>
