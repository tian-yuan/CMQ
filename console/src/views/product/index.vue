<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="120px">
      <el-form-item label="ProductKey">
        <el-input v-model="form.ProductKey"/>
        <el-button type="primary" @click="onSubmit">查询</el-button>
      </el-form-item>
    </el-form>
    <json-editor ref="jsonEditor" v-model="form.value"/>
  </div>
</template>

<script>
import JsonEditor from '@/components/JsonEditor'
import { getProductInfo } from '@/api/product'

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
      var params = { 'ProductKey': this.form.ProductKey }
      getProductInfo(params).then(response => {
        this.form.value = response.ProductInfo
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

