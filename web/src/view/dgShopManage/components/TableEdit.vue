<template>
  <el-dialog :title="title" :visible.sync="dialogFormVisible" width="500px" @close="close">
    <!-- 信息展示栏 -->
    <el-form ref="formRef" :model="formData" label-width="80px" size="mini">
      <el-form-item label="店铺头像" prop="shop_url">
        <el-image :src="formData.shop_url" style="width: 40px;"></el-image>
      </el-form-item>
      <el-form-item label="店铺ID号" prop="shop_id">
        <el-input readonly v-model.trim="formData.shop_id" autocomplete="off"></el-input>
      </el-form-item>
      <el-form-item label="店铺名称" prop="shop_name">
        <el-input readonly v-model.trim="formData.shop_name" autocomplete="off"></el-input>
      </el-form-item>
      <el-form-item label="手机号码" prop="phone_number">
        <el-input readonly v-model.trim="formData.phone_number" autocomplete="off"></el-input>
      </el-form-item>
      <el-form-item label="操作选项">
        <el-select @change="operateChange" v-model="operateSelectVal" filterable placeholder="请选择">
          <el-option-group v-for="group in operateOptions" :key="group.label" :label="group.label">
            <el-option
              v-for="item in group.options"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            ></el-option>
          </el-option-group>
        </el-select>
      </el-form-item>
    </el-form>
    <!-- 操作选项 -->

    <div slot="footer" class="dialog-footer">
      <el-button @click="close">取 消</el-button>
      <el-button type="primary" @click="save">确 定</el-button>
    </div>
  </el-dialog>
</template>

<script>
import { operateShop } from "@/api/daigou/shopManage";
import { respCode } from "@/utils/errorCode";

const EnumOperate = {
  None: "",
  FreezeShop: 0,
  UnFreezeShop: 1,
  CloseShop: 2,
  OpenShop: 3,
  UpgradeShop: 4,
};

export default {
  name: "TableEdit",
  data() {
    return {
      formData: {
        shop_id: "",
        shop_name: "",
        phone_number: "",
        shop_url: "",
      },
      //对话框标题
      title: "",
      //是否隐藏
      dialogFormVisible: false,
      // 操作选项
      operateOptions: [
        {
          label: "最近操作",
          options: [],
        },
        {
          label: "所有操作",
          options: [
            {
              value: EnumOperate.FreezeShop,
              label: "冻结店铺",
            },
            {
              value: EnumOperate.UnFreezeShop,
              label: "解冻店铺",
            },
            {
              value: EnumOperate.CloseShop,
              label: "关闭店铺",
            },
            {
              value: EnumOperate.OpenShop,
              label: "开启店铺",
            },
            {
              value: EnumOperate.UpgradeShop,
              label: "升级店铺",
            },
          ],
        },
      ],
      operateSelectVal: EnumOperate.None, //操作选项当前选择值
    };
  },
  created() {},
  methods: {
    operateChange(val) {
      console.error(this.$options.data());
      const newOptions = [this.operateOptions[1].options[val]];
      for (const v of this.operateOptions[0].options) {
        if (v.value == val) {
          continue;
        }
        newOptions.push(v);
        if (newOptions.length >= 3) {
          break;
        }
      }
      this.operateOptions[0].options = newOptions;
    },

    showEdit(row) {
      // this.title = "选择要进行的操作操作";
      this.formData = Object.assign({}, row);
      this.dialogFormVisible = true;
    },
    close() {
      this.$refs["formRef"].resetFields();
      this.formData = this.$options.data().formData;
      this.dialogFormVisible = false;
    },
    checkOperate() {
      switch (this.operateSelectVal) {
        case EnumOperate.FreezeShop:
          if (this.formData.is_enable == 0) {
            return "店铺当前已经是冻结状态";
          }
          break;
        case EnumOperate.UnFreezeShop:
          if (this.formData.is_enable == 1) {
            return "店铺当前已经是解冻状态";
          }
          break;
        default:
          return "未定义的操作";
      }
      return "";
    },

    async save() {
      // 提交前进行检查
      const msg1 = this.checkOperate();
      if (msg1 !== "") {
        this.$message.warning(msg1);
        return;
      }
      let request = {};
      switch (this.operateSelectVal) {
        case EnumOperate.FreezeShop:
          request = {
            shop_id: this.formData.shop_id,
            operate: EnumOperate.FreezeShop,
          };
          break;
        case EnumOperate.UnFreezeShop:
          request = {
            shop_id: this.formData.shop_id,
            operate: EnumOperate.UnFreezeShop,
          };
          break;
        default:
          console.error("未定义的操作");
          return;
      }

      const { code, msg } = await operateShop(request);
      if (code != respCode.SUCCESS) {
        this.$message.error(msg);
        return;
      }
      this.$message.success(msg);
      this.$refs["formRef"].resetFields();
      this.dialogFormVisible = false;
      this.$emit("fetchData");
      this.formData = this.$options.data().formData;
    },
  },
};
</script>
