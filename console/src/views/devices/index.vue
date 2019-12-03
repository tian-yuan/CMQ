<template>
  <div class="app-container">
    <div class="filter-container">
      <el-input v-model="listQuery.Keyword" :placeholder="$t('Keyword')" style="width: 200px;" class="filter-item" @keyup.enter.native="handleFilter" />
      <el-input v-model="listQuery.ProductKey" :placeholder="$t('ProductKey')" style="width: 200px;" class="filter-item" @keyup.enter.native="handleFilter" />
      <el-button v-waves class="filter-item" style="margin-left: 10px;margin-bottom: 10px;" type="primary" icon="el-icon-search" @click="handleFilter">
        {{ $t('table.search') }}
      </el-button>
    </div>

    <el-table
      v-loading="listLoading"
      :key="tableKey"
      :data="list"
      border
      fit
      highlight-current-row
      style="width: 100%;"
    >
      <el-table-column :label="$t('DeviceName')" min-width="150px">
        <template slot-scope="scope">
          <span class="link-type" @click="handleDeviceDetail(scope.row)">{{ scope.row.DeviceName }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('ProductKey')" width="150px">
        <template slot-scope="scope">
          <span>{{ scope.row.ProductKey }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('DeviceSecret')" width="150px">
        <template slot-scope="scope">
          <span>{{ scope.row.DeviceSecret }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('State')" width="150px">
        <template slot-scope="scope">
          <span>{{ scope.row.State }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('CreateAt')" width="150px">
        <template slot-scope="scope">
          <span>{{ scope.row.CreateAt | parseTime('{y}-{m}-{d} {h}:{i}') }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('LastActiveAt')" width="150px">
        <template slot-scope="scope">
          <span>{{ scope.row.LastActiveAt | parseTime('{y}-{m}-{d} {h}:{i}') }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('UpdateAt')" width="150px">
        <template slot-scope="scope">
          <span>{{ scope.row.UpdateAt | parseTime('{y}-{m}-{d} {h}:{i}') }}</span>
        </template>
      </el-table-column>
    </el-table>

    <pagination v-show="total>0" :total="total" :page.sync="listQuery.page" :limit.sync="listQuery.limit" @pagination="getList" />

    <el-dialog :visible.sync="dialogVisible" title="Device Detail">
      <el-form ref="temp" :model="temp" label-width="200px">
        <el-form-item label="DeviceName">
          <el-input v-model="temp.DeviceName"/>
        </el-form-item>
      </el-form>

      <span slot="footer" class="dialog-footer">
        <el-button type="primary" @click="dialogVisible = false">{{ $t('confirm') }}</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>

import { fetchDeviceList } from '@/api/devices'
import waves from '@/directive/waves' // Waves directive
import { parseTime } from '@/utils'
import Pagination from '@/components/Pagination' // Secondary package based on el-pagination
import JsonEditor from '@/components/JsonEditor'

export default {
  name: 'ComplexTable',
  components: { Pagination, JsonEditor },
  directives: { waves },
  filters: {
    statusFilter(status) {
      const statusMap = {
        published: 'success',
        draft: 'info',
        deleted: 'danger'
      }
      return statusMap[status]
    },
    parseTime: function(timestamp, format) {
      return parseTime(timestamp, format)
    }
  },
  data() {
    return {
      tableKey: 0,
      list: null,
      total: 0,
      listLoading: true,
      listQuery: {
        page: 1,
        limit: 20,
        ProductKey: undefined,
        Keyword: undefined
      },
      sortOptions: [{ label: 'ID Ascending', key: '+id' }, { label: 'ID Descending', key: '-id' }],
      rules: {
        title: [{ required: true, message: 'title is required', trigger: 'blur' }]
      },
      temp: {
        DeviceName: ''
      },
      dialogVisible: false
    }
  },
  created() {
    this.getList()
  },
  methods: {
    getList() {
      this.listLoading = true
      var params = {
        ProductKey: this.listQuery.ProductKey,
        Keyword: this.listQuery.Keyword,
        Offset: (this.listQuery.page - 1) * this.listQuery.limit,
        Limit: this.listQuery.limit
      }
      fetchDeviceList(params).then(response => {
        this.list = response.DeviceInfoList
        console.log('device info list : ', this.list)
        this.total = response.TotalCount

        // Just to simulate the time of the request
        setTimeout(() => {
          this.listLoading = false
        }, 1.5 * 1000)
      })
    },
    handleFilter() {
      this.listQuery.page = 1
      this.getList()
    },
    handleDeviceDetail(row) {
      var params = { 'ProductKey': row.ProductKey, 'DeviceName': row.DeviceName }
      console.log('params info : ', params)
      this.temp.DeviceName = row.DeviceName
      this.dialogVisible = true
    },
    formatJson(filterVal, jsonData) {
      return jsonData.map(v => filterVal.map(j => {
        if (j === 'timestamp') {
          return parseTime(v[j])
        } else {
          return v[j]
        }
      }))
    }
  }
}
</script>
