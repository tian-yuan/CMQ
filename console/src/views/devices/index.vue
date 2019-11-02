<template>
  <div class="app-container">
    <div class="filter-container">
      <el-input v-model="listQuery.DeviceId" :placeholder="$t('table.DeviceId')" style="width: 200px;" class="filter-item" @keyup.enter.native="handleFilter" />
      <el-date-picker v-model="listQuery.Start" type="datetime" placeholder="Please pick a date"/>
      <el-date-picker v-model="listQuery.End" type="datetime" placeholder="Please pick a date"/>
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
      <el-table-column :label="$t('table.MessageId')" min-width="150px">
        <template slot-scope="scope">
          <span class="link-type" @click="handleMessageDetail(scope.row)">{{ scope.row.MessageId }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('table.MessageType')" width="150px">
        <template slot-scope="scope">
          <span>{{ scope.row.MessageType }}</span>
        </template>
      </el-table-column>

      <el-table-column :label="$t('table.Timestamp')" width="150px">
        <template slot-scope="scope">
          <span>{{ scope.row.Timestamp | parseTime('{y}-{m}-{d} {h}:{i}') }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('table.Domain')" min-width="150px">
        <template slot-scope="scope">
          <span>{{ scope.row.Domain }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('table.Platform')" width="110px">
        <template slot-scope="scope">
          <span>{{ scope.row.Platform }}</span>
        </template>
      </el-table-column>
    </el-table>

    <pagination v-show="total>0" :total="total" :page.sync="listQuery.page" :limit.sync="listQuery.limit" @pagination="getList" />
    <el-dialog :visible.sync="dialogVisible" title="Message Detail">
      <el-form ref="temp" :model="temp" label-width="120px">
        <el-form-item label="Message Id">
          <el-input v-model="temp.MessageId"/>
        </el-form-item>
      </el-form>
      <json-editor ref="jsonEditor" v-model="temp.Content"/>

      <span slot="footer" class="dialog-footer">
        <el-button type="primary" @click="dialogVisible = false">{{ $t('table.confirm') }}</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import { getPushListByDeviceId } from '@/api/device'
import { getMessageInfo } from '@/api/message'
import waves from '@/directive/waves' // Waves directive
import { parseTime } from '@/utils'
import Pagination from '@/components/Pagination' // Secondary package based on el-pagination
import JsonEditor from '@/components/JsonEditor'

export default {
  name: 'ComplexTable',
  components: { Pagination, JsonEditor },
  directives: { waves },
  filters: {
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
        Content: undefined,
        Platform: undefined,
        Channel: undefined,
        MessageType: undefined,
        title: undefined,
        type: undefined
      },
      temp: {
        MessageId: '',
        Content: ''
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
        DeviceId: this.listQuery.DeviceId,
        Start: new Date(this.listQuery.Start).getTime(),
        End: new Date(this.listQuery.End).getTime(),
        Offset: (this.listQuery.page - 1) * this.listQuery.limit,
        Limit: this.listQuery.limit
      }
      getPushListByDeviceId(params).then(response => {
        this.list = response.data.BriefMessageInfoList
        console.log('brief message info list : ', this.list)
        this.total = response.data.TotalCount

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
    handleMessageDetail(row) {
      var params = { 'MessageId': row.MessageId }
      getMessageInfo(params).then(response => {
        this.temp.MessageId = row.MessageId
        this.temp.Content = JSON.parse(response.data)
        this.dialogVisible = true
      })
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
