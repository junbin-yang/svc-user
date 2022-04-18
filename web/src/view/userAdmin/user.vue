<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="searchForm" :inline="true" :model="searchInfo">
        <el-form-item label="账号">
          <el-input v-model="searchInfo.account" placeholder="账号" />
        </el-form-item>
        <el-form-item label="手机号">
          <el-input v-model="searchInfo.mobile" placeholder="手机号" />
        </el-form-item>
        <el-form-item label="邮箱">
          <el-input v-model="searchInfo.email" placeholder="邮箱" />
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
      <el-table :data="tableData" @sort-change="sortChange" @selection-change="handleSelectionChange">
        <el-table-column
          type="selection"
          width="55"
        />
        <el-table-column align="left" label="id" min-width="60" prop="ID" sortable="custom" />
        <el-table-column align="left" label="账号" min-width="100" prop="Account" sortable="custom" />
        <el-table-column align="left" label="昵称" min-width="100" prop="NickName" sortable="custom" />
        <el-table-column align="left" label="手机号" min-width="120" prop="Mobile" sortable="custom" />
        <el-table-column align="left" label="邮箱" min-width="120" prop="Email" sortable="custom" />
        <el-table-column align="left" label="状态" min-width="80" prop="Status" sortable="custom">
          <template #default="scope">
            <div>
              {{ statusFiletr(scope.row.Status) }}
            </div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="最近登录时间" min-width="120" prop="LoginedAt" sortable="custom">
          <template #default="scope">
            <div>
              {{ momentTime(scope.row.LoginedAt) }}
            </div>
          </template>
        </el-table-column>

        <el-table-column align="left" fixed="right" label="操作" width="200">
          <template #default="scope">
            <el-button
              icon="edit"
              size="small"
              type="text"
              @click="editUserFunc(scope.row)"
            >编辑</el-button>
            <el-button type="text" icon="magic-stick" size="small" @click="resetPasswordFunc(scope.row)">重置密码</el-button>
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
      <el-form ref="userForm" :model="userform" :rules="rules" label-width="80px">
        <el-form-item label="账号" prop="Account">
          <el-input v-model="userform.Account" autocomplete="off" disabled />
        </el-form-item>
        <el-form-item label="昵称" prop="NickName">
          <el-input v-model="userform.NickName" autocomplete="off" />
        </el-form-item>
        <el-form-item label="手机号" prop="Mobile">
          <el-input v-model="userform.Mobile" autocomplete="off" />
        </el-form-item>
        <el-form-item label="邮箱" prop="Email">
          <el-input v-model="userform.Email" autocomplete="off" />
        </el-form-item>
        <el-form-item label="姓名" prop="Info.RealName">
          <el-input v-model="userform.Info.RealName" autocomplete="off" />
        </el-form-item>
        <el-form-item label="年龄" prop="Info.Age">
          <el-input v-model="userform.Info.Age" autocomplete="off" />
        </el-form-item>
        <el-form-item label="性别" prop="Info.Sex">
          <el-select v-model="userform.Info.Sex" placeholder="请选择" style="width:100%">
            <el-option
              v-for="item in SexOptions"
              :key="item.value"
              :label="`${item.label}(${item.value})`"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="民族" prop="Info.Nation">
          <el-input v-model="userform.Info.Nation" autocomplete="off" />
        </el-form-item>
        <el-form-item label="身份证" prop="Info.IdCard">
          <el-input v-model="userform.Info.IdCard" autocomplete="off" />
        </el-form-item>
        <el-form-item label="地址" prop="Info.Address">
          <el-input v-model="userform.Info.Address" autocomplete="off" />
        </el-form-item>
        <el-form-item label="状态" prop="Status">
          <el-select v-model="userform.Status" placeholder="请选择" style="width:100%">
            <el-option
              v-for="item in StatusOptions"
              :key="item.value"
              :label="`${item.label}(${item.value})`"
              :value="item.value"
            />
          </el-select>
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
  name: 'User',
}
</script>

<script setup>
import { getUserList, updateUser, resetPassword } from '@/api/user'
import { toSQLLine } from '@/utils/stringFun'
import { ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import moment from 'moment'

const momentTime = (value) => {
  return moment(value).format('YYYY-MM-DD hh:mm:ss')
}

const userform = ref({
  Account: '',
  Email: '',
  Status: '',
  Mobile: '',
  NickName: '',
  Info: {},
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

const SexOptions = ref([
  {
    value: 0,
    label: '女',
    type: 'success'
  },
  {
    value: 1,
    label: '男',
    type: 'danger'
  }
])

const type = ref('')
const rules = ref({
  Account: [{ required: true, message: '请输入账号', trigger: 'blur' }],
  NickName: [{ required: true, message: '请输入昵称', trigger: 'blur' }],
  Mobile: [
    { required: true, message: '请输入手机号', trigger: 'blur' }
  ],
  Email: [
    { required: true, message: '请输入邮箱地址', trigger: 'blur' }
  ],
  Status: [
    { required: true, message: '请选择状态', trigger: 'blur' }
  ]
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
  if (searchInfo.value.account) {
    searchObj[searchInfo.value.account] = ['account']
  }
  if (searchInfo.value.mobile) {
    searchObj[searchInfo.value.mobile] = ['mobile']
  }
  if (searchInfo.value.email) {
    searchObj[searchInfo.value.email] = ['email']
  }
  if (searchInfo.value.status) {
    searchObj[searchInfo.value.status] = ['status']
  }
  const table = await getUserList({ page: page.value, limit: pageSize.value, search: searchObj })
  if (table.code === 200) {
    tableData.value = table.data.docs
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.limit
  }
}

getTableData()

// 弹窗相关
const userForm = ref(null)
const initForm = () => {
  userForm.value.resetFields()
  userform.value = {
    Account: '',
    Email: '',
    Status: '',
    Mobile: '',
    NickName: '',
    Info: {},
  }
}

const dialogTitle = ref('新增用户')
const dialogFormVisible = ref(false)
const openUserDialog = (key) => {
  switch (key) {
    case 'addUser':
      dialogTitle.value = '新增用户'
      break
    case 'edit':
      dialogTitle.value = '编辑用户'
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

const editUserFunc = async(row) => {
  userform.value = row
  openUserDialog('edit')
}

const enterDialog = async() => {
  userForm.value.validate(async valid => {
    if (valid) {
      switch (type.value) {
        /*
        case 'addUser':
          {
            const res = await createUser(userform.value)
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
          */
        case 'edit':
          {
            const res = await updateUser(userform.value)
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

const resetPasswordFunc = (row) => {
  ElMessageBox.confirm(
    '是否将此用户密码重置?',
    '警告',
    {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning',
    }
  ).then(async() => {
    const res = await resetPassword({
      email: row.Email,
    })
    if (res.code === 200) {
      /*
      ElMessage({
        type: 'success',
        message: '修改成功',
      })
      */
      ElMessageBox.alert(res.data, '新密码')
    } else {
      ElMessage({
        type: 'error',
        message: res.message,
      })
    }
  })
}

</script>

<style scoped lang="scss">
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
