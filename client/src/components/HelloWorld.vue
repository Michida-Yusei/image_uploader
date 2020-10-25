<template>
  <div class="hello">
    <vue-dropzone ref="myVueDropzone" id="dropzone" :options="dropzoneOptions"></vue-dropzone>
  </div>
</template>

<script>
  // vue2-dropzone と vue2-dropzone用のcssをimport
  import vue2Dropzone from 'vue2-dropzone'
  import 'vue2-dropzone/dist/vue2Dropzone.min.css'

  export default {
    name: 'HelloWorld',
    data: function () {
      return {
        dropzoneOptions: {
          url: `http://localhost:8888/images`,
          method: 'post'
        }
      }
    },
    components: {
      vueDropzone: vue2Dropzone
    }
  }
</script>
