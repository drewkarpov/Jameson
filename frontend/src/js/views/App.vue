<template>
    <b-container class="jameson-container">
        <b-row>
            <b-col class="mt-3 mb-3">
                <h1>Image Comparator</h1>

                <router-link to="/test" class="router-link">Compare</router-link>
                <router-link :to="voidZonesPath" class="router-link">Define voidzones</router-link>


            </b-col>
        </b-row>

        <router-view @change-container-id="handleChangeContainerId"/>
    </b-container>
</template>
<script>
export default {
  name: 'App',

  data() {
    return {
      containerId: null
    }
  },

  computed: {
    voidZonesPath() {
      const path = {
        name: 'container'
      }

      if(this.containerId) {
        path.params = {
          containerIdParam: this.containerId
        }
      }

      return path
    }
  },

  methods: {
    handleChangeContainerId(id) {
      console.log(id)
      this.containerId =id
    }
  }

}
</script>


