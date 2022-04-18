<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="searchForm" :inline="true" :model="searchInfo">
        <el-form-item label="名称">
          <el-input v-model="searchInfo.name" placeholder="名称" />
        </el-form-item>
        <el-form-item label="描述">
          <el-input v-model="searchInfo.description" placeholder="描述" />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchInfo.status" clearable placeholder="请选择">
            <el-option
              v-for="item in StatusOptions"
              :key="item.value"
              :label="`${item.label}(${item.value})`"
              :value="item.value?'1':'0'"
            />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="small" type="primary" icon="plus" @click="openDialog('addApp')">新增</el-button>
      </div>
      <el-table :data="tableData" @sort-change="sortChange" @selection-change="handleSelectionChange">
        <el-table-column
          type="selection"
          width="55"
        />
        <el-table-column align="left" label="id" min-width="60" prop="ID" sortable="custom" />
        <el-table-column align="left" label="名称" min-width="100" prop="Name" sortable="custom" />
        <el-table-column align="left" label="AppKey" min-width="120" prop="AppKey" sortable="custom" />
        <el-table-column align="left" label="AppSecret" min-width="200" prop="AppSecret" sortable="custom" />
        <el-table-column align="left" label="状态" min-width="80" prop="Status" sortable="custom">
          <template #default="scope">
            <div>
              {{ statusFiletr(scope.row.Status) }}
            </div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="备注" min-width="150" prop="Remark" sortable="custom" />

        <el-table-column align="left" fixed="right" label="操作" width="200">
          <template #default="scope">
            <el-button
              size="small"
              icon="document"
              type="text"
              @click="toDetile(scope.row)"
            >角色详情</el-button>
            <el-button
              icon="edit"
              size="small"
              type="text"
              @click="editAppFunc(scope.row)"
            >编辑</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          layout="total, sizes, prev, pager, next, jumper"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>

    </div>

    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" :title="dialogTitle">
      <el-form ref="apiForm" :model="form" :rules="rules" label-width="90px">
        <el-form-item label="名称" prop="Name">
          <el-input v-model="form.Name" autocomplete="off" />
        </el-form-item>
        <el-form-item v-if="type==='edit'" label="AppKey" prop="AppKey">
          <el-input v-model="form.AppKey" autocomplete="off" disabled />
        </el-form-item>
        <el-form-item v-if="type==='edit'" label="AppSecret" prop="AppSecret">
          <el-input v-model="form.AppSecret" autocomplete="off" />
        </el-form-item>
        <el-form-item v-if="type==='edit'" label="状态" prop="Status">
          <el-select v-model="form.Status" placeholder="请选择" style="width:100%">
            <el-option
              v-for="item in StatusOptions"
              :key="item.value"
              :label="`${item.label}(${item.value})`"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="描述" prop="Remark">
          <el-input v-model="form.Remark" autocomplete="off" />
        </el-form-item>
        <el-form-item v-if="type==='addApp'" label="验证码" prop="validateCode">
          <div class="vPicBox">
            <el-input
              v-model="form.validateCode"
              style="width: 60%"
              placeholder="请输入验证码"
              show-password
            />
            <div class="vPic">
              <img
                v-if="picPath"
                :src="picPath"
                alt="请输入验证码"
                @click="getVerify()"
              >
            </div>
          </div>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'App',
}
</script>

<script setup>
import { getAppList, getAppById, updateApp, createApp } from '@/api/app'
import { toSQLLine } from '@/utils/stringFun'
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import { captcha } from '@/api/user'
import { useRouter } from 'vue-router'
const router = useRouter()

const form = ref({
  Name: '',
  AppKey: '',
  AppSecret: '',
  Status: '',
  Remark: '',
  validateCode: '',
  validateCodeId: '',
})

const statusFiletr = (value) => {
  const target = StatusOptions.value.filter(item => item.value === value)[0]
  return target && `${target.label}`
}

const StatusOptions = ref([
  {
    value: true,
    label: '启用',
    type: 'success'
  },
  {
    value: false,
    label: '禁用',
    type: 'danger'
  }
])

const type = ref('')
const rules = ref({
  Name: [{ required: true, message: '请输入名称', trigger: 'blur' }],
  AppKey: [
    { required: true, message: '请输入AppKey', trigger: 'blur' }
  ],
  AppSecret: [
    { required: true, message: '请输入AppSecret', trigger: 'blur' }
  ],
  Status: [
    { required: true, message: '请选择状态', trigger: 'blur' }
  ],
  validateCode: [{ required: true, message: '请输入验证码', trigger: 'blur' }],
})

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

const onReset = () => {
  searchInfo.value = {}
}
// 搜索

const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 排序
const sortChange = ({ prop, order }) => {
  if (prop) {
    if (prop === 'ID') {
      prop = 'id'
    }
    searchInfo.value.orderKey = toSQLLine(prop)
    searchInfo.value.desc = order === 'descending'
  }
  getTableData()
}

// 查询
const getTableData = async() => {
  const searchObj = {}
  if (searchInfo.value.name) {
    searchObj[searchInfo.value.name] = ['name']
  }
  if (searchInfo.value.description) {
    searchObj[searchInfo.value.description] = ['remark']
  }
  if (searchInfo.value.status) {
    searchObj[searchInfo.value.status] = ['status']
  }
  const table = await getAppList({ page: page.value, limit: pageSize.value, search: searchObj })
  if (table.code === 200) {
    tableData.value = table.data.docs
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.limit
  }
}

getTableData()

// 弹窗相关
const apiForm = ref(null)
const initForm = () => {
  apiForm.value.resetFields()
  form.value = {
    Name: '',
    AppKey: '',
    AppSecret: '',
    Status: '',
    Remark: '',
    validateCode: '',
    validateCodeId: '',
  }
}

const dialogTitle = ref('新增App')
const dialogFormVisible = ref(false)
const openDialog = (key) => {
  switch (key) {
    case 'addApi':
      dialogTitle.value = '新增App'
      break
    case 'edit':
      dialogTitle.value = '编辑App'
      break
    default:
      break
  }
  type.value = key
  dialogFormVisible.value = true
}
const closeDialog = () => {
  initForm()
  dialogFormVisible.value = false
}

const editAppFunc = async(row) => {
  const res = await getAppById({ id: row.ID })
  form.value = res.data.app
  openDialog('edit')
}

const enterDialog = async() => {
  apiForm.value.validate(async valid => {
    if (valid) {
      switch (type.value) {
        case 'addApp':
          {
            const res = await createApp(form.value)
            if (res.code === 200) {
              ElMessage({
                type: 'success',
                message: '添加成功',
                showClose: true
              })
            }
            getTableData()
            closeDialog()
          }

          break
        case 'edit':
          {
            const res = await updateApp(form.value)
            if (res.code === 200) {
              ElMessage({
                type: 'success',
                message: '编辑成功',
                showClose: true
              })
            }
            getTableData()
            closeDialog()
          }
          break
        default:
          // eslint-disable-next-line no-lone-blocks
          {
            ElMessage({
              type: 'error',
              message: '未知操作',
              showClose: true
            })
          }
          break
      }
    }
  })
}

const picPath = ref('')
// 获取验证码
const getVerify = () => {
  captcha({}).then((ele) => {
    picPath.value = ele.data.b64s
    form.value.validateCodeId = ele.data.key
  })
}
getVerify()

const toDetile = (row) => {
  router.push({
    name: 'appDetail',
    params: {
      id: row.AppKey,
    },
  })
}

</script>

<style scoped lang="scss">
.vPicBox{
    display:flex;
    justify-content:space-between;
    width:100%;
}
.vPic {
    width: 33%;
    height: 38px;
    background: #ccc;
    img {
        width: 100%;
        height: 100%;
        vertical-align: middle;
    }
}
.button-box {
  padding: 10px 20px;
  .el-button {
    float: right;
  }
}
.warning {
  color: #dc143c;
}
</style>
