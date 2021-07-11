<template>

  <el-form :model="this" label-width="100px">
    <el-form-item label="启用">
      <el-switch v-model="enable"/>
    </el-form-item>
    <el-form-item label="主机">
      <el-input v-model="hostname"/>
    </el-form-item>
    <el-form-item label="端口">
      <el-input v-model="port" type="number"/>
    </el-form-item>
    <el-form-item label="SSL">
      <el-switch v-model="ssl"/>
    </el-form-item>
    <el-form-item label="用户名">
      <el-input v-model="username"/>
    </el-form-item>
    <el-form-item label="密码">
      <el-input v-model="password" type="password" show-password/>
    </el-form-item>
    <el-form-item label="发件人">
      <el-input v-model="fromName"/>
    </el-form-item>
    <el-form-item label="发件邮箱">
      <el-input v-model="fromEmail"/>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="onSubmit">保存</el-button>
      <el-button type="primary" @click="onTest">测试</el-button>
    </el-form-item>
  </el-form>

  <el-dialog title="发送测试邮件" v-model="dialog" :close-on-click-modal="false">
    <el-form ref="dialog" :model="dialogData" :rules="dialogRules" label-width="120px">
      <el-form-item label="邮箱" prop="email">
        <el-input v-model="dialogData.email" placeholder="请输入名邮箱地址"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onDialogSubmit">发送</el-button>
        <el-button @click="dialog = false">取消</el-button>
      </el-form-item>
    </el-form>
  </el-dialog>

</template>

<script lang="ts">
import { defineComponent } from 'vue';
import {updateSMTP, querySMTP, testSMTP} from "@/api";

export default defineComponent({
  name: 'SMTP',
  components: {},
  mounted() {
    this.onRefresh()
  },
  data() {
    return {
      enable: false,
      hostname: null,
      port: null,
      ssl: false,
      username: null,
      password: null,
      fromName: null,
      fromEmail: null,

      dialog: false,
      dialogData: {},
      dialogRules: {
        email: [
          {required: true, message: '邮箱不能为空', trigger: 'blur'},
        ],
      },
    }
  },
  methods: {
    onRefresh() {
      querySMTP()
          .then(data => {
            data = data.data
            this.enable = data.enable;
            this.hostname = data.hostname;
            this.port = data.port;
            this.ssl = data.ssl;
            this.username = data.username;
            this.password = data.password;
            this.fromName = data.fromName;
            this.fromEmail = data.fromEmail;
          })
    },

    onSubmit() {
      updateSMTP(
          this.enable,
          this.hostname,
          this.port,
          this.ssl,
          this.username,
          this.password,
          this.fromName,
          this.fromEmail,
      )
          .then(() => {
            this.$message.success("更新成功");
            this.onRefresh()
          })
    },
    onTest() {
      this.dialog = true
    },

    onDialogSubmit() {
      let from: any = this.$refs['dialog'];
      from.validate((valid: boolean) => {
        if (!valid) return
        let dialogData = this.dialogData as any
        testSMTP(dialogData.email)
            .then(() => {
              this.$message.success("发送成功");
              this.dialog = false
            })
      })
    },
  }
});
</script>

<style lang="less">
</style>
