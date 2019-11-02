<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="120px">
      <el-form-item label="消息ID">
        <el-input v-model="form.messageId"/>
        <el-button type="primary" @click="onSubmit">查询</el-button>
      </el-form-item>
    </el-form>
    <json-editor ref="jsonEditor" v-model="form.value"/>
  </div>
</template>

<script>
import JsonEditor from '@/components/JsonEditor'
import { getMessageInfo } from '@/api/message'

export default {
  components: { JsonEditor },
  data() {
    return {
      form: {
        messageId: '',
        value: ''
      }
    }
  },
  methods: {
    onSubmit() {
      var params = { 'MessageId': this.form.messageId }
      getMessageInfo(params).then(response => {
        this.form.value = JSON.parse(response.MessageInfo)
      })
    }
  }
}
</script>

<style scoped>
.line{
  text-align: center;
}
.editor-container{
  position: relative;
  height: 100%
}
</style>

