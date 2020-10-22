<template>
  <div>
    <el-divider content-position="center">请求参数</el-divider>
    <el-form
      ref="formReqRef"
      :model="operateReq"
      :rules="argsRules"
      label-width="80px"
      size="mini"
      label-position="left"
    >
      <el-form-item label="店铺码" prop="shop_code">
        <el-input
          v-model.trim="operateReq.shop_code"
          autocomplete="off"
          placeholder="店铺码格式`BU2ACA`"
        ></el-input>
      </el-form-item>
      <el-form-item label="类型" prop="gen_type">
        <el-input readonly v-model.trim="genType" autocomplete="off"></el-input>
      </el-form-item>
      <el-form-item label="绑定店铺" prop="shop_id">
        <el-input readonly v-model.trim="operateReq.shop_id" autocomplete="off"></el-input>
      </el-form-item>
      <el-form-item label="有效期(s)" prop="valid_time">
        <el-input readonly v-model.trim="operateReq.valid_time" autocomplete="off"></el-input>
      </el-form-item>
    </el-form>
    <el-button type="primary" @click="handelGetShopCode">店铺码</el-button>
  </div>
</template>

<script>
import {
  proxyGetShopCode,
  proxyRespCode,
  proxyUseShopCode,
} from "@/api/daigou/proxyApi";
import { reactive, computed, toRefs } from "@vue/composition-api";

export default {
  name: "OpenShop",
  props: [],

  setup(props, { refs, root }) {
    const checkShopCode = (rule, value, callback) => {
      const reg = /^[A-Z0-9]+$/;
      if (!value) {
        return callback(new Error("店铺码不能为空"));
      } else if (
        typeof value !== "string" ||
        value.length != 6 ||
        !reg.test(value)
      ) {
        return callback(new Error("请输入正确店铺码"));
      } else {
        callback();
      }
    };

    const { $message, $notify } = root;

    const state = reactive({
      // 旧店铺码
      old_shop_code: "",
      //表单请求参数
      operateReq: {},
      // 表单校验规则
      argsRules: {
        shop_code: [{ validator: checkShopCode, trigger: "blur" }],
      },
      rolwName: {
        1: "用户",
        2: "个体商户",
        3: "推广员",
        4: "代购",
        5: "总代",
      },

      genType: computed(() => {
        console.error(state.operateReq);
        if (state.operateReq === undefined) {
          return "";
        }
        return state.operateReq.gen_type === 0 ? "系统生成" : "店铺生成";
      }),
    });

    const checkArgs = () => {
      let result = false;
      refs.formReqRef.validate((valid) => {
        if (!valid) {
          $message.error("参数不正确");
          result = false;
        } else {
          result = true;
        }
      });
      return result;
    };
    //执行操作
    const operate = async ({ user_id }) => {
      if (!checkArgs() || state.operateReq.shop_code != state.old_shop_code) {
        $message.error(`请输入有效的店铺码`);
        return;
      }

      const { code, data, msg } = await proxyUseShopCode({
        user_id,
        shop_code: state.operateReq.shop_code,
      });

      console.error({ func: "operate", data, code, msg });

      if (code != proxyRespCode.SUCCESS) {
        $message.error(`开通店铺失败,${msg}`);
        return;
      }
      return data;
    };
    //结果操作处理
    const showOperate = ({ data, user_id }) => {
      console.error({ func: "showResponse", data });
      const role = state.rolwName[data.role];
      let message = `用户${user_id}开通店铺成功,店铺号为${data.shop_id},角色变为${role}`;
      if (data.bind_shop_id != 0) {
        message += `,成功绑定店铺${data.bind_shop_id}`;
      }
      $notify({
        title: "成功",
        message: message,
        type: "success",
        duration: 0,
      });
    };

    //获取数据
    const fetchData = async () => {
      const { code, msg, data } = await proxyGetShopCode({
        gen_type: 0,
        old_shop_code: state.old_shop_code,
      });

      if (code != proxyRespCode.SUCCESS) {
        $message.error(`代理请求店铺码错误,${msg}`);
        return;
      }
      console.error({ data, func: "handelGetShopCode" });
      state.operateReq = data;
      state.old_shop_code = data.shop_code;
    };

    const handelGetShopCode = () => {
      fetchData();
    };

    fetchData();
    return {
      ...toRefs(state),
      checkArgs,
      operate,
      showOperate,
      fetchData,
      handelGetShopCode,
    };
  },
};
</script>
<style lang="scss">
</style>