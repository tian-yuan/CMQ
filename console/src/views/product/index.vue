<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="120px">
      <el-form-item label="ProductKey">
        <el-col :span="11">
          <el-input v-model="form.ProductKey"/>
        </el-col>
        <el-col :span="11">
          <el-button type="primary" @click="onSubmit">查询</el-button>
        </el-col>
      </el-form-item>
      <el-form-item label="ProductName">
        <el-input v-model="form.ProductName"/>
      </el-form-item>
      <el-form-item label="ProductDetail">
        <el-input v-model="form.Description"/>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onProductCreate">创建</el-button>
      </el-form-item>
    </el-form>
    <json-editor ref="jsonEditor" v-model="form.value"/>
  </div>
</template>

<script>
import JsonEditor from '@/components/JsonEditor'
import { getProductInfo, createProduct } from '@/api/product'
import ElForm from '../../../node_modules/element-ui/packages/form/src/form'
import ElFormItem from '../../../node_modules/element-ui/packages/form/src/form-item'

export default {
  components: {
    ElFormItem,
    ElForm,
    JsonEditor },
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
    },
    onProductCreate() {
      console.log('create product.')
      var params = { 'ProductName': this.form.ProductName, 'Description': this.form.Description }
      createProduct(params).then(response => {
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

