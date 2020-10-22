/* 公共引入,勿随意修改,修改时需经过确认 */
import '@/config/permission'
import { VabQueryForm } from '@/layouts'
import Vab from '@/utils/vab'
import Vue from 'vue'
import '../utils/vab'
import './bus'
import './clipboard2'
import './composition-api'
import './element'
import './mavon-editor'
import './vue-particle-line'
import './vue-quill-editor'
import './vue-simple-uploader'
Vue.use(VabQueryForm)
Vue.use(Vab)
