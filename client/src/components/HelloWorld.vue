<template>
  <div class="hello">

    <vue-dropzone ref="myVueDropzone" id="dropzone" :options="dropzoneOptions"
                  v-on:vdropzone-sending="sendingEvent"
                  v-on:vdropzone-removed-file="removeEvent"
    ></vue-dropzone>

  </div>
</template>

<script>
  import vue2Dropzone from 'vue2-dropzone'
  import 'vue2-dropzone/dist/vue2Dropzone.min.css'
  import axios from 'axios'

  export default {
    name: 'HelloWorld',
    data: function () {
      return {
        dropzoneOptions: {
          url: `http://localhost:8888/images`,
          method: 'post',
          addRemoveLinks: 'true',
          dictDefaultMessage: "<i class='fa fa-cloud-upload'></i>UPLOAD ME"
        }
      }
    },
    components: {
      vueDropzone: vue2Dropzone
    },
    methods: {
      // formデータとして fileに付けられた任意のuuidを付加
      sendingEvent: function (file, xhr, formData) {
        formData.append('uuid', file.upload.uuid)
      },
      removeEvent: function (file) {
        axios
        .delete(`http://localhost:8888/images/${file.upload.uuid}`)
        .then(res => {
          console.log(res.data)
        }).catch(err => {
          console.error(err)
        })
      }
    }
  }
</script>
<style>
  @import url("https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css");
</style>
