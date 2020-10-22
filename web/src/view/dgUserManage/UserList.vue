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
          <el-form-item prop="nickName">
            <el-input :style="{width: '300px'}" v-model="queryForm.nickName " placeholder="微信昵称" />
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

      <el-table-column
        show-overflow-tooltip
        label="用户ID"
        prop="user_id"
        min-width="88"
        sortable="custom"
        fixed
      ></el-table-column>

      <el-table-column show-overflow-tooltip label="微信昵称" prop="nick_name" min-width="79" fixed></el-table-column>

      <el-table-column show-overflow-tooltip label="用户头像" prop="avatar_url" min-width="79">
        <template v-slot:default="scope">
          <el-image
            v-if="imgShow"
            :preview-src-list="imageList"
            :src="scope.row.avatar_url"
            style="width: 40px;"
          ></el-image>
        </template>
      </el-table-column>

      <el-table-column show-overflow-tooltip label="性别" prop="gender" min-width="52" fixed>
        <template v-slot:default="scope">
          <el-tag :type="scope.row.gender | filterGender">{{ scope.row.gender | filterGenderStr}}</el-tag>
        </template>
      </el-table-column>

      <el-table-column show-overflow-tooltip label="手机号码" prop="phone_number" min-width="108"></el-table-column>

      <el-table-column show-overflow-tooltip label="店铺ID号" prop="shop_id" min-width="79"></el-table-column>

      <el-table-column show-overflow-tooltip label="绑定店铺" prop="bind_shop_id" min-width="79"></el-table-column>

      <el-table-column show-overflow-tooltip label="城市" prop="city" min-width="79"></el-table-column>

      <el-table-column show-overflow-tooltip label="OpenID" prop="open_id" min-width="79"></el-table-column>

      <el-table-column show-overflow-tooltip label="公众号" prop="gzh_open_id" min-width="79"></el-table-column>

      <el-table-column
        show-overflow-tooltip
        label="注册时间"
        prop="created_at"
        min-width="152"
        sortable="custom"
      ></el-table-column>
      <el-table-column
        show-overflow-tooltip
        label="最后登录"
        prop="updated_at"
        min-width="152"
        sortable="custom"
      ></el-table-column>

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
    <table-operate ref="tableOperateRef" @fetchData="fetchData"></table-operate>
  </div>
</template>

<script>
import { deleteUserInfoByIds, findUserInfoList } from "@/api/daigou/userManage";
import TableOperate from "./components/TableOperate";
import { respCode } from "@/utils/errorCode";
import { isBlank, isPhone } from "@/utils/validate";

import { reactive, toRefs } from "@vue/composition-api";

const dbOrderSwitch = {
  ascending: "asc",
  descending: "desc",
};
export default {
  name: "UserList",
  components: {
    TableOperate,
  },

  filters: {
    //是否推荐状态
    filterGender(status) {
      const statusMap = {
        0: "info",
        1: "",
        2: "danger",
      };
      return statusMap[status];
    },
    filterGenderStr(status) {
      const statusMap = {
        0: "未知",
        1: "男",
        2: "女",
      };
      return statusMap[status];
    },
  },
  setup(props, { refs, root }) {
    const validPhoneNumber = (rule, value, callback) => {
      // console.error({ name: "validPhoneNumber", rule, value, callback });
      if (!isBlank(value) && !isPhone(value)) {
        callback(new Error("请输入正确的手机号"));
      } else {
        callback();
      }
    };
    //实例属性
    const { $message, $baseConfirm, $baseMessage, $copyText } = root; // <--- setup() 中不可以使用 this 访问当前组件实例, 可以通过 setup 的第二个参数 context 来访问 vue2.x API 中实例上的属性。
    // 此处先省略 data 部分...
    const state = reactive({
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
        nickName: "",
        phoneNumber: "",
      },
      //上次查询的条件
      oldCondition: {},
      //查询条件校验规则
      formRules: {
        phoneNumber: [
          {
            validator: validPhoneNumber,
            trigger: "blur",
          },
        ],
        nickName: [],
      },
      //粘贴板
      clipboard: {},
      //重新获取数据
      bRegainList: false,
      //上次排序标记
      lastOrders: "user_id asc",
    });

    //获取数据
    const fetchData = async (orders = state.lastOrders) => {
      console.error("fetchData");
      state.listLoading = true;
      const { data, code, msg } = await findUserInfoList({
        page: state.queryForm.page,
        pageSize: state.queryForm.pageSize,
        orders,
      });
      if (code != respCode.SUCCESS) {
        $message.error(msg);
        return;
      }
      state.list = data.list;
      //用户头像大图预览
      const imageList = [];
      data.list.forEach((item, index) => {
        imageList.push(item.avatar_url);
        state.list[index].index = index;
      });
      state.imageList = imageList;
      state.total = data.total;
      state.listLoading = false;
    };
    //排序变化
    const tableSortChange = ({ prop, order }) => {
      const currentOrder = `${prop} ${dbOrderSwitch[order]}`;
      if (state.lastOrders === currentOrder || order === null) {
        console.log("已经排序");
        return;
      }
      fetchData(currentOrder);
      state.lastOrders = currentOrder;
    };

    //选中的列改变
    const setSelectRows = (val) => {
      state.selectRows = val;
    };
    const handleOper = (row) => {
      refs["tableOperateRef"].showOpera(row);
    };

    const handleDelete = (row) => {
      if (row.id) {
        $baseConfirm("你确定要删除当前项吗", null, async () => {
          const { msg } = await deleteUserInfoByIds({
            ids: row.id,
          });
          $baseMessage(msg, "success");
          fetchData();
        });
      } else {
        if (state.selectRows.length > 0) {
          const ids = state.selectRows.map((item) => item.id).join();
          $baseConfirm("你确定要删除选中项吗", null, async () => {
            const { msg } = await deleteUserInfoByIds({
              ids: ids,
            });
            $baseMessage(msg, "success");
            fetchData();
          });
        } else {
          $baseMessage("未选中任何行", "error");
          return false;
        }
      }
    };
    //每页个数变化
    const handleSizeChange = (val) => {
      state.queryForm.pageSize = val;
      fetchData();
    };

    //双击单元格
    const handleCellDbclick = (row, column, cell, event) => {
      $copyText(cell.textContent).then(
        (e) => {
          $message.success(`复制${column.property}成功:${e.text}`);
        },
        (e) => {
          console.error({
            action: e.action,
            trigger: e.trigger,
          });
          $message.error("复制失败");
        }
      );
    };
    //当前页变化时
    const handleCurrentPageChange = (val) => {
      state.queryForm.page = val;
      fetchData();
    };
    //点击查询按钮
    const handleQuery = (formName) => {
      refs[formName].validate(async (valid) => {
        if (!valid) {
          console.error("请检查查询参数");
          return false;
        }
        //检查查询条件是否全部为空
        if (
          isBlank(state.queryForm.phoneNumber) &&
          isBlank(state.queryForm.nickName)
        ) {
          $message.warning("查询条件为空");
          return;
        }
        //检查查询条件是否更改
        if (
          state.oldCondition.phoneNumber === state.queryForm.phoneNumber &&
          state.oldCondition.nickName === state.queryForm.nickName
        ) {
          console.log("查询条件未更改");
          return;
        }
        state.queryForm.page = 1;

        const { data, code, msg } = await findUserInfoList({
          page: state.queryForm.page,
          pageSize: state.queryForm.pageSize,
          phoneNumber: state.queryForm.phoneNumber,
          nickName: state.queryForm.nickName,
        });

        if (code != respCode.SUCCESS) {
          $message.error(msg);
          return;
        }
        state.list = data.list;
        //用户头像大图预览
        const imageList = [];
        data.list.forEach((item, index) => {
          imageList.push(item.user_url);
          state.list[index].index = index;
        });
        state.imageList = imageList;
        state.total = data.total;

        //记录上次查询条件
        state.oldCondition = {
          phoneNumber: state.queryForm.phoneNumber,
          nickName: state.queryForm.nickName,
        };
        state.bRegainList = true;
      });
    };
    //重置
    const resetQuery = (formName) => {
      refs[formName].resetFields();
      if (state.bRegainList) {
        fetchData();
        state.oldCondition = {};
        state.bRegainList = false;
      }
    };

    fetchData();

    return {
      ...toRefs(state),
      resetQuery,
      handleQuery,
      handleCurrentPageChange,
      handleCellDbclick,
      handleSizeChange,
      handleDelete,
      handleOper,
      setSelectRows,
      tableSortChange,
      fetchData,
    };
  },
};
</script>
