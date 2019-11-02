<template>
  <div class="app-container">
    <div class="filter-container">
      <el-input v-model="listQuery.DeviceId" :placeholder="$t('table.DeviceId')" style="width: 200px;" class="filter-item" @keyup.enter.native="handleFilter" />
      <el-button v-waves class="filter-item" style="margin-left: 10px;margin-bottom: 10px;" type="primary" icon="el-icon-search" @click="handleFilter">
        {{ $t('table.search') }}
      </el-button>
    </div>

    <el-table
      v-loading="listLoading"
      :data="list"
      border
      fit
      highlight-current-row
      style="width: 100%;"
    >
      <el-table-column :label="$t('table.Item')" width="150px">
        <template slot-scope="scope">
          <span>{{ scope.row.Item }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('table.Value')" min-width="150px">
        <template slot-scope="scope">
          <span>{{ scope.row.Value }}</span>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import { getDeviceInfo } from '@/api/device'
import waves from '@/directive/waves' // Waves directive

export default {
  name: 'ComplexTable',
  directives: { waves },
  data() {
    return {
      list: null,
      listLoading: true,
      listQuery: {
        DeviceId: undefined
      }
    }
  },
  created() {
    this.getList()
  },
  methods: {
    getList() {
      this.listLoading = true
      var params = { 'DeviceId': this.listQuery.DeviceId }
      getDeviceInfo(params).then(response => {
        var deviceInfo = response.data
        console.log('brief device info : ', deviceInfo)
        if (deviceInfo) {
          this.list = [
            { Item: 'DeviceId', Value: deviceInfo.DeviceId },
            { Item: 'Domain', Value: deviceInfo.Domain },
            { Item: 'Platform', Value: deviceInfo.Platform },
            { Item: 'SubPlatform', Value: deviceInfo.SubPlatform },
            { Item: 'Token', Value: deviceInfo.Token }
          ]
        }
        // Just to simulate the time of the request
        setTimeout(() => {
          this.listLoading = false
        }, 1 * 1000)
      })
    },
    handleFilter() {
      this.getList()
    }
  }
}
</script>
