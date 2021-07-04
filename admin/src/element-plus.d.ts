import {ElMessage, ElMessageBox} from 'element-plus';
declare module '@vue/runtime-core' {
  interface ComponentCustomProperties {
    $message: typeof ElMessage,
    $alert: typeof ElMessageBox.alert,
    $prompt: typeof ElMessageBox.prompt,
    $confirm: typeof ElMessageBox.confirm,
  }
}