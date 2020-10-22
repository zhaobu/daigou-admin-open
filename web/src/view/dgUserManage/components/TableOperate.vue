<template>
  <el-dialog :title="title" :visible.sync="dialogFormVisible" width="500px" @close="close">
    <!-- 基本信息 -->
    <el-divider content-position="center">用户信息</el-divider>
    <el-form
      ref="formBaseRef"
      :model="formBaseData"
      label-width="80px"
      size="mini"
      label-position="left"
    >
      <el-form-item label="用户头像" prop="avatar_url">
        <el-image :src="formBaseData.avatar_url" style="width: 40px;"></el-image>
      </el-form-item>
      <el-form-item label="用户ID号" prop="user_id">
        <el-input readonly v-model.trim="formBaseData.user_id" autocomplete="off"></el-input>
      </el-form-item>
      <el-form-item label="微信昵称" prop="nick_name">
        <el-input readonly v-model.trim="formBaseData.nick_name" autocomplete="off"></el-input>
      </el-form-item>
      <el-form-item label="手机号码" prop="phone_number">
        <el-input readonly v-model.trim="formBaseData.phone_number" autocomplete="off"></el-input>
      </el-form-item>
      <el-form-item label="店铺ID号" prop="shop_id">
        <el-input readonly v-model.trim="formBaseData.shop_id" autocomplete="off"></el-input>
      </el-form-item>
    </el-form>

    <!-- 操作选项 -->
    <el-divider></el-divider>
    <el-form label-width="80px" size="mini" label-position="left">
      <el-form-item label="操作选项">
        <el-select @change="operateChange" v-model="operateSelectVal" filterable placeholder="请选择">
          <el-option-group v-for="group of operateOptions" :key="group.label" :label="group.label">
            <el-option
              v-for="item of group.options"
              :key="item.value"
              :label="item.label"
              :value="item.value"
              :disabled="item.disabled"
            ></el-option>
          </el-option-group>
        </el-select>
      </el-form-item>
    </el-form>

    <!-- 请求参数信息 -->
    <template v-if="isShowReqArg">
      <open-shop v-if="isShow(0)" ref="operateRef"></open-shop>
      <close-shop v-else-if="isShow(1)" ref="operateRef"></close-shop>
    </template>

    <el-divider></el-divider>

    <div v-show="isShowReqArg" slot="footer">
      <el-button @click="close">取 消</el-button>
      <el-button type="primary" @click="save">确 定</el-button>
    </div>
  </el-dialog>
</template>

<script>
import OpenShop from "./OpenShop";
import CloseShop from "./CloseShop";
import { reactive, computed, toRefs } from "@vue/composition-api";

const EnumOperate = {
  None: undefined,
  OpenShop: 0,
  CloseShop: 1,
};

export default {
  name: "TableOperate",
  components: {
    OpenShop,
    CloseShop,
  },

  setup(props, { refs, root }) {
    const { $options, $emit } = root;

    const initState = {
      // 基本信息
      formBaseData: {
        user_id: "",
        nick_name: "",
        phone_number: "",
        avatar_url: "",
      },

      // 当前操作的参数信息
      // operateReq: {},

      //对话框标题
      title: "",
      //是否隐藏
      dialogFormVisible: false,
      // 操作选项
      operateOptions: {
        // latest: {
        //   label: "最近操作",
        //   options: {},
        // },
        all: {
          label: "所有操作",
          options: {
            0: {
              value: EnumOperate.OpenShop,
              label: "开通店铺",
              disabled: false,
            },
            1: {
              value: EnumOperate.CloseShop,
              label: "关闭店铺",
              disabled: true,
            },
          },
        },
      },
      operateSelectVal: EnumOperate.None, //操作选项当前选择值
      isShowReqArg: false, //是否显示请求参数

      isShow: computed(() => {
        return function (operate) {
          return operate == state.operateSelectVal;
        };
      }),
    };

    const state = reactive(Object.assign({}, initState));

    const operateChange = () => {
      switch (state.operateSelectVal) {
        case EnumOperate.OpenShop:
          state.isShowReqArg = true;
          break;
        default:
          return;
      }
    };

    const showOpera = (row) => {
      // state.title = "选择要进行的操作操作";
      state.formBaseData = Object.assign({}, row);
      state.dialogFormVisible = true;
      const options = state.operateOptions.all.options;
      //禁用不能操作的
      if (state.formBaseData.shop_id !== 0) {
        options[EnumOperate.OpenShop].disabled = true;
      }
    };

    const close = () => {
      resetData();
    };

    const save = async () => {
      //调用子组件的方法进行校验
      if (!refs.operateRef) {
        return;
      }

      //执行操作
      const data = await refs.operateRef.operate({
        user_id: state.formBaseData.user_id,
      });
      //展示处理结果
      refs.operateRef.showOperate({
        data,
        user_id: state.formBaseData.user_id,
      });
      // savelastOper(state.operateSelectVal);
      resetData();
      $emit("fetchData");
    };

    //记录最近操作的选项
    // const savelastOper = (operID) => {
    //   const newOption = state.operateOptions.all.options[operID];
    //   const latestOptions = state.operateOptions.latest.options;
    //   const newOptions = { newOption };
    //   for (const key in latestOptions) {
    //     newOptions.key = latestOptions[key];
    //     //最多记录最近三个
    //     if (Object.keys(newOptions).length >= 3) {
    //       break;
    //     }
    //   }
    //   state.operateOptions.latest.options = newOptions;
    // };

    const resetData = () => {
      console.error({ $options });
      refs["formBaseRef"].resetFields();
      Object.assign(state.formBaseData, initState.formBaseData);
      Object.assign(state.operateOptions, initState.operateOptions);
      // state.formBaseData = $options.data().formBaseData;
      // state.operateOptions.all = $options.data().operateOptions.all;
      state.dialogFormVisible = false;
      state.operateSelectVal = EnumOperate.None;
      state.isShowReqArg = false;
    };

    return {
      ...toRefs(state),
      operateChange,
      showOpera,
      close,
      save,
      resetData,
    };
  },
};
</script>
