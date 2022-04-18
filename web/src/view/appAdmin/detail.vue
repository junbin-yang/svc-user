<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo">
        <el-form-item label="角色名">
          <el-input v-model="searchInfo.name" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="描述">
          <el-input v-model="searchInfo.desc" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="启用状态" prop="status">
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
        <el-button size="small" type="primary" icon="plus" @click="openDialog('create')">创建应用角色</el-button>
      </div>
      <el-table
        ref="multipleTable"
        :data="tableData"
        style="width: 100%"
        tooltip-effect="dark"
        row-key="ID"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="id" prop="ID" min-width="60" />
        <el-table-column align="left" label="角色名" prop="Name" min-width="100" />
        <el-table-column align="left" label="创建日期" min-width="120">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>

        <el-table-column align="left" label="描述" prop="Remark" min-width="150" />

        <el-table-column align="left" label="启用状态" min-width="80" prop="Status">
          <template #default="scope">
            <div>
              {{ statusFiletr(scope.row.Status) }}
            </div>
          </template>
        </el-table-column>

        <el-table-column align="left" label="用户数" prop="UserCount" min-width="80" />

        <el-table-column align="left" label="按钮组">
          <template #default="scope">
            <el-button
              icon="edit"
              size="small"
              type="text"
              @click="editRoleFunc(scope.row)"
            >编辑</el-button>
            <el-button
              icon="setting"
              size="small"
              type="text"
              @click="bindUserFunc(scope.row)"
            >绑定用户</el-button>
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
      <div v-if="type =='bindUser'" style="text-align: center">
        <el-transfer
          v-model="transferData.rightValue"
          style="text-align: left; display: inline-block"
          filterable
          :titles="['未绑定', '已绑定']"
          :button-texts="['解绑', '绑定']"
          :format="{
            noChecked: '${total}',
            hasChecked: '${checked}/${total}',
          }"
          :data="transferData.data"
          @change="handleTransferChange"
        >
          <template #default="{ option }">
            <span>{{ option.label }}</span>
          </template>
        </el-transfer>
      </div>
      <el-form v-if="type !=='bindUser'" ref="dialogForm" :model="formData" :rules="rules" size="medium" label-width="110px">
        <el-form-item label="名称" prop="Name">
          <el-input v-model="formData.Name" autocomplete="off" />
        </el-form-item>
        <el-form-item label="描述" prop="Remark">
          <el-input v-model="formData.Remark" autocomplete="off" />
        </el-form-item>
        <el-form-item label="启用状态" prop="Status">
          <el-select v-model="formData.Status" placeholder="请选择" style="width:100%">
            <el-option
              v-for="item in StatusOptions"
              :key="item.value"
              :label="`${item.label}(${item.value})`"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="角色权限" prop="Auth">
          <el-tree
            ref="treeRef"
            :data="treeData"
            show-checkbox
            default-expand-all
            node-key="ID"
            highlight-current
            :props="defaultProps"
            :default-checked-keys="isChecked"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button v-if="type !=='bindUser'" size="small" type="primary" @click="enterDialog">确 定</el-button>
          <el-button v-if="type ==='bindUser'" size="small" type="primary" @click="enterTransferDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'RolesDetail'
}
</script>

<script setup>
import { getAppRoleList, updateRole, createRole, bindRoleUser } from '@/api/app'
import { getUserList } from '@/api/user'
import { ref } from 'vue'
import { useRoute } from 'vue-router'
import { ElMessage } from 'element-plus'
import { formatDate } from '@/utils/format'
const route = useRoute()

const transferData = ref({
  data: [],
  rightValue: [],
  roleId: null,
  userIds: []
})
const handleTransferChange = async(value, direction, movedKeys) => {
  console.log(value, direction, movedKeys)
  /*
  value.forEach(item => {
    transferData.value.userIds.push(item)
  })
  */
  transferData.value.userIds = value
  /*
  const res = await bindRoleUser({ AppId: String(route.params.id), RoleId: transferData.value.roleId, UserId: value })
  if (res.code === 200) {
    ElMessage({
      type: 'success',
      message: res.message
    })
  }
  */
}
const getAllUserData = async() => {
  const data = []
  const res = await getUserList({ page: 1, limit: 999999 })
  if (res.code === 200) {
    res.data.docs.forEach(itme => {
      data.push({
        key: itme.ID,
        label: `${itme.NickName} - [ ${itme.Account} ] `,
        disabled: !itme.Status,
      })
    })
  }
  transferData.value.data = data
}
getAllUserData()

const treeRef = ref('')
const isChecked = ref([])
const getChecked = () => {
  const out = {}
  const ckNodes = treeRef.value.getCheckedNodes()
  ckNodes.forEach(item => {
    if (item.IsLeaf === true) {
      out[item.PCode] |= item.Value
    }
  })
  console.info(out)
  return out
}

const defaultProps = {
  children: 'Children',
  label: 'Name',
  isLeaf: 'IsLeaf',
}

const treeData = ref([])

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

const formData = ref({
  ID: 0,
  AppKey: '',
  Name: '',
  Status: false,
  Remark: '',
  RoleAuthoritys: {},
  Sequence: '',
  Other: {}
})
const rules = ref({
  Name: [
    {
      required: true,
      message: '请输入角色名',
      trigger: 'blur'
    }
  ],
  Status: [
    {
      required: true,
      message: '请选择状态',
      trigger: 'blur'
    }
  ]
})

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({ appId: String(route.params.id) })
const onReset = () => {
  searchInfo.value = { appId: String(route.params.id) }
}

// 条件搜索前端看此方法
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

// 查询
const getTableData = async() => {
  const searchObj = {}
  searchObj[searchInfo.value.appId] = ['app_key']
  if (searchInfo.value.name) {
    searchObj[searchInfo.value.name] = ['name']
  }
  if (searchInfo.value.desc) {
    searchObj[searchInfo.value.desc] = ['remark']
  }
  if (searchInfo.value.status) {
    searchObj[searchInfo.value.status] = ['status']
  }
  const table = await getAppRoleList(
    { page: page.value, limit: pageSize.value, search: searchObj }
  )
  if (table.code === 200) {
    tableData.value = table.data.roles.docs
    total.value = table.data.roles.total
    page.value = table.data.roles.page
    pageSize.value = table.data.roles.limit
    treeData.value = table.data.roleAuthTree
  }
}

getTableData()

const type = ref('')
const dialogFormVisible = ref(false)

const closeDialog = () => {
  isChecked.value = []
  dialogFormVisible.value = false
  formData.value = {
    ID: 0,
    AppKey: '',
    Name: '',
    Status: false,
    Remark: '',
    RoleAuthoritys: {},
    Sequence: '',
    Other: {}
  }
}

const dialogForm = ref(null)
const enterDialog = async() => {
  dialogForm.value.validate(async valid => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createRole({ AppId: String(route.params.id), Name: formData.value.Name, Status: formData.value.Status, Remark: formData.value.Remark, Permissions: getChecked() })
        break
      case 'edit':
        res = await updateRole({ AppId: String(route.params.id), Id: formData.value.ID, Name: formData.value.Name, Status: formData.value.Status, Remark: formData.value.Remark, Permissions: getChecked() })
        break
      case 'bindUser':
        // res = await bindRoleUser({ AppId: String(route.params.id), RoleId: transferData.value.roleId, UserId: transferData.value.userIds })
        console.info({ AppId: String(route.params.id), RoleId: transferData.value.roleId, UserId: transferData.value.userIds })
        res.code = 200
        break
      default:
        break
    }
    if (res.code === 200) {
      ElMessage({
        type: 'success',
        message: res.message
      })
      closeDialog()
      getTableData()
    }
  })
}
const enterTransferDialog = async() => {
  const res = await bindRoleUser({ AppId: String(route.params.id), RoleId: transferData.value.roleId, UserId: transferData.value.userIds })
  if (res.code === 200) {
    ElMessage({
      type: 'success',
      message: res.message
    })
    closeDialog()
    getTableData()
  }
}

const dialogTitle = ref('')
const openDialog = (key) => {
  switch (key) {
    case 'create':
      dialogTitle.value = '创建角色'
      break
    case 'edit':
      dialogTitle.value = '编辑角色'
      break
    case 'bindUser':
      dialogTitle.value = '绑定用户'
      break
    default:
      dialogTitle.value = '创建角色'
      break
  }
  type.value = key
  dialogFormVisible.value = true
}

const editRoleFunc = async(row) => {
  formData.value = row
  isChecked.value = formData.value.Other.checkedNode
  openDialog('edit')
}

const bindUserFunc = async(row) => {
  transferData.value.rightValue = row.Other.assUserIds
  transferData.value.roleId = row.ID
  openDialog('bindUser')
}

</script>

<style>
</style>
