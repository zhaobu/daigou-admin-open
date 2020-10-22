<template>
  <div class="table-container">
    <vab-query-form>
      <vab-query-form-left-panel :style="{width: '100%'}">
        <!-- <el-button icon="el-icon-plus" type="primary" @click="handleAdd">添加</el-button> -->
        <el-button icon="el-icon-delete" type="danger" @click="handleDelete">批量删除</el-button>
        <!-- 搜索区域 -->
        <el-form
          ref="queryFormRef"
          :model="queryForm"
          size="medium"
          :inline="true"
          :rules="formRules"
        >
          <el-form-item prop="phoneNumber">
            <el-input :style="{width: '120px'}" v-model="queryForm.phoneNumber" placeholder="手机号码" />
          </el-form-item>
          <el-form-item prop="shopName">
            <el-input :style="{width: '300px'}" v-model="queryForm.shopName " placeholder="店铺名称" />
          </el-form-item>
          <el-form-item>
            <el-button icon="el-icon-search" type="primary" @click="handleQuery('queryFormRef')">查询</el-button>
            <el-button icon="el-icon-refresh" @click="resetQuery('queryFormRef')">重置</el-button>
          </el-form-item>
        </el-form>
      </vab-query-form-left-panel>
    </vab-query-form>
    <!-- 数据表格 -->
    <el-table
      ref="tableSortRef"
      v-loading="listLoading"
      :data="list"
      :element-loading-text="elementLoadingText"
      @selection-change="setSelectRows"
      @sort-change="tableSortChange"
      @cell-dblclick="handleCellDbclick"
      max-height="700"
      border
      highlight-current-row
      row-key="id"
    >
      <el-table-column show-overflow-tooltip type="selection" label="选择" min-width="40"></el-table-column>
      <el-table-column show-overflow-tooltip label="序号" min-width="52" prop="id" fixed>
        <!-- <template v-slot:default="scope">{{ scope.$index + 1 }}</template> -->
      </el-table-column>

      <el-table-column
        show-overflow-tooltip
        label="店铺ID"
        prop="shop_id"
        min-width="88"
        sortable="custom"
        fixed
      ></el-table-column>

      <el-table-column show-overflow-tooltip label="店铺名称" prop="shop_name" min-width="79" fixed></el-table-column>

      <el-table-column show-overflow-tooltip label="手机号码" prop="phone_number" min-width="108" fixed></el-table-column>

      <el-table-column show-overflow-tooltip label="店铺头像" prop="shop_url" min-width="79">
        <template v-slot:default="scope">
          <el-image
            v-if="imgShow"
            :preview-src-list="imageList"
            :src="scope.row.shop_url"
            style="width: 40px;"
          ></el-image>
        </template>
      </el-table-column>

      <el-table-column
        show-overflow-tooltip
        label="粉丝数"
        prop="shop_fans_count"
        min-width="88"
        sortable="custom"
      ></el-table-column>

      <el-table-column
        show-overflow-tooltip
        label="创建时间"
        prop="created_at"
        min-width="152"
        sortable="custom"
      ></el-table-column>

      <el-table-column show-overflow-tooltip label="店铺说明" prop="shop_description" min-width="79"></el-table-column>

      <el-table-column show-overflow-tooltip label="微信号" prop="wechat_number" min-width="64"></el-table-column>

      <el-table-column show-overflow-tooltip label="是否推荐" prop="is_recommend" min-width="79">
        <template v-slot:default="scope">
          <el-tooltip :content="scope.row.is_recommend | statusRecommendStr" placement="top">
            <el-tag :type="scope.row.is_recommend | statusRecommend">{{ scope.row.is_recommend }}</el-tag>
          </el-tooltip>
        </template>
      </el-table-column>

      <el-table-column show-overflow-tooltip label="商铺状态" prop="is_enable" min-width="79">
        <template v-slot:default="scope">
          <el-tooltip :content="scope.row.is_enable | statusShopStr" placement="top">
            <el-tag :type="scope.row.is_enable | statusShop">{{ scope.row.is_enable}}</el-tag>
          </el-tooltip>
        </template>
      </el-table-column>

      <el-table-column show-overflow-tooltip label="分类信息" prop="category_info" min-width="79"></el-table-column>

      <el-table-column
        show-overflow-tooltip
        label="滚动信息"
        prop="mainpage_scroll_info"
        min-width="79"
      ></el-table-column>

      <el-table-column show-overflow-tooltip label="用户ID" prop="user_id" min-width="76"></el-table-column>

      <el-table-column show-overflow-tooltip label="图片水印配置" prop="shop_watermark" min-width="120"></el-table-column>

      <el-table-column show-overflow-tooltip label="收款二维码" prop="qr_code_url" min-width="95"></el-table-column>

      <el-table-column show-overflow-tooltip label="操作" min-width="55" fixed="right">
        <template v-slot:default="scope">
          <el-button type="text" @click="handleOper(scope.row)">操作</el-button>
        </template>
      </el-table-column>
    </el-table>
    <!-- 分页底部 -->
    <el-pagination
      :background="background"
      :current-page="queryForm.page"
      :layout="layout"
      :page-size="queryForm.pageSize"
      :total="total"
      @current-change="handleCurrentPageChange"
      @size-change="handleSizeChange"
    ></el-pagination>
    <table-edit ref="editRef" @fetchData="fetchData"></table-edit>
  </div>
</template>

<script>
import { deleteShopInfoByIds, findShopInfoList } from "@/api/daigou/shopManage";
import TableEdit from "./components/TableEdit";
// import EditableCell from "./components/EditableCell";
import { respCode } from "@/utils/errorCode";
import { isBlank, isPhone } from "@/utils/validate";

const dbOrderSwitch = {
  ascending: "asc",
  descending: "desc",
};
export default {
  name: "ShopList",
  components: {
    TableEdit,
    // EditableCell,
  },
  filters: {
    //是否推荐状态
    statusRecommend(status) {
      const statusMap = {
        0: "gray", //不推荐
        1: "success", //推荐
      };
      return statusMap[status];
    },
    statusRecommendStr(status) {
      const statusMap = {
        0: "不推荐", //不推荐
        1: "推荐", //推荐
      };
      return statusMap[status];
    },
    //商铺状态
    statusShop(status) {
      const statusMap = {
        0: "gray", //不启用
        1: "success", //启用
      };
      return statusMap[status];
    },
    statusShopStr(status) {
      const statusMap = {
        0: "不启用", //不启用
        1: "启用", //启用
      };
      return statusMap[status];
    },
  },
  data() {
    const validPhoneNumber = (rule, value, callback) => {
      // console.error({ name: "validPhoneNumber", rule, value, callback });
      if (!isBlank(value) && !isPhone(value)) {
        callback(new Error("请输入正确的手机号"));
      } else {
        callback();
      }
    };
    return {
      //是否显示图像
      imgShow: true,
      //列表信息
      list: [],
      //大图预览
      imageList: [],
      // 加载数据标识
      listLoading: true,
      //底部分页显示控制
      layout: "total, sizes, prev, pager, next, jumper",
      //总共条数
      total: 0,
      background: true,
      //选中的所有列
      selectRows: "",
      //加载提示信息
      elementLoadingText: "正在加载...",
      // 查询数据
      queryForm: {
        page: 1,
        pageSize: 20,
        shopName: "",
        phoneNumber: "",
      },
      //上次查询的条件
      oldCondition: {},
      //查询条件校验规则
      formRules: {
        phoneNumber: [{ validator: validPhoneNumber, trigger: "blur" }],
        shopName: [],
      },
      //粘贴板
      clipboard: {},
      //重新获取数据
      bRegainList: false,
      //上次排序标记
      lastOrders: "user_id asc",
    };
  },
  created() {
    this.fetchData();
  },
  beforeDestroy() {},
  mounted() {},
  methods: {
    //排序变化
    tableSortChange({ column, prop, order }) {
      console.error({ column, prop, order });

      const currentOrder = `${prop} ${dbOrderSwitch[order]}`;
      if (this.lastOrders === currentOrder || order === null) {
        console.log("已经排序");
        return;
      }
      this.fetchData(currentOrder);
      this.lastOrders = currentOrder;
    },
    //选中的列改变
    setSelectRows(val) {
      this.selectRows = val;
    },
    handleOper(row) {
      this.$refs["editRef"].showEdit(row);
    },
    handleDelete(row) {
      if (row.id) {
        this.$baseConfirm("你确定要删除当前项吗", null, async () => {
          const { msg } = await deleteShopInfoByIds({ ids: row.id });
          this.$baseMessage(msg, "success");
          this.fetchData();
        });
      } else {
        if (this.selectRows.length > 0) {
          const ids = this.selectRows.map((item) => item.id).join();
          this.$baseConfirm("你确定要删除选中项吗", null, async () => {
            const { msg } = await deleteShopInfoByIds({ ids: ids });
            this.$baseMessage(msg, "success");
            this.fetchData();
          });
        } else {
          this.$baseMessage("未选中任何行", "error");
          return false;
        }
      }
    },
    //每页个数变化
    handleSizeChange(val) {
      this.queryForm.pageSize = val;
      this.fetchData();
    },

    //双击单元格
    handleCellDbclick(row, column, cell, event) {
      console.error({ func: "handleCellDbclick", row, column, cell, event });
      this.$copyText(cell.textContent).then(
        (e) => {
          this.$message.success(`复制${column.property}成功:${e.text}`);
        },
        (e) => {
          console.error({ action: e.action, trigger: e.trigger });
          this.$message.error("复制失败");
        }
      );
    },
    //当前页变化时
    handleCurrentPageChange(val) {
      this.queryForm.page = val;
      this.fetchData();
    },
    //点击查询按钮
    handleQuery(formName) {
      this.$refs[formName].validate(async (valid) => {
        if (!valid) {
          console.error("请检查查询参数");
          return false;
        }
        //检查查询条件是否全部为空
        if (
          isBlank(this.queryForm.phoneNumber) &&
          isBlank(this.queryForm.shopName)
        ) {
          this.$message.warning("查询条件为空");
          return;
        }
        //检查查询条件是否更改
        if (
          this.oldCondition.phoneNumber === this.queryForm.phoneNumber &&
          this.oldCondition.shopName === this.queryForm.shopName
        ) {
          console.log("查询条件未更改");
          return;
        }
        this.queryForm.page = 1;

        const { data, code, msg } = await findShopInfoList({
          page: this.queryForm.page,
          pageSize: this.queryForm.pageSize,
          phoneNumber: this.queryForm.phoneNumber,
          shopName: this.queryForm.shopName,
        });

        if (code != respCode.SUCCESS) {
          this.$message.error(msg);
          return;
        }
        this.list = data.list;
        //店铺头像大图预览
        const imageList = [];
        data.list.forEach((item, index) => {
          imageList.push(item.shop_url);
          this.list[index].index = index;
        });
        this.imageList = imageList;
        this.total = data.total;

        //记录上次查询条件
        this.oldCondition = {
          phoneNumber: this.queryForm.phoneNumber,
          shopName: this.queryForm.shopName,
        };
        this.bRegainList = true;
      });
    },
    //重置
    resetQuery(formName) {
      this.$refs[formName].resetFields();
      if (this.bRegainList) {
        this.fetchData();
        this.oldCondition = {};
        this.bRegainList = false;
      }
    },
    //获取数据
    async fetchData(orders = this.lastOrders) {
      console.error("fetchData");
      if (orders === undefined) {
        orders = ["shop_id asc"];
      }
      this.listLoading = true;
      const { data, code, msg } = await findShopInfoList({
        page: this.queryForm.page,
        pageSize: this.queryForm.pageSize,
        orders,
      });
      if (code != respCode.SUCCESS) {
        this.$message.error(msg);
        return;
      }
      this.list = data.list;
      //店铺头像大图预览
      const imageList = [];
      data.list.forEach((item, index) => {
        imageList.push(item.shop_url);
        this.list[index].index = index;
      });
      this.imageList = imageList;
      this.total = data.total;
      this.listLoading = false;
    },
  },
};
</script>
