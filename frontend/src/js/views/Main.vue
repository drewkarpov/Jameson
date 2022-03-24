<template>
  <div>
    <b-row>
      <b-col>
        <b-form-input v-model="testId" @keyup.enter="update()" placeholder="Test Id"></b-form-input>
      </b-col>
      <b-col>
        <b-button block variant="primary" @click="update">Update</b-button>
      </b-col>
    </b-row>

    <b-row>
      <b-col class="mt-3 mb-3">
        <h2>Results diff →
          <percentage :value="percentage"></percentage>
        </h2>
      </b-col>
      <b-col class="mt-3 mb-3">
        <h2>Name →
          <description :value="testName"></description>
        </h2>
      </b-col>
      <b-col class="mt-3 mb-3">
        <h2>Container_Id →
          <description :value="containerId"></description>
        </h2>
      </b-col>
    </b-row>


    <b-row>
      <b-col>
        <b-tabs content-class="mt-3" fill v-model="activeTab">
          <b-tab title="Reference">
            <b-img fluid-grow rounded :src="images.reference"></b-img>
          </b-tab>
          <b-tab title="Candidate">
            <b-img fluid-grow rounded :src="images.candidate"></b-img>
          </b-tab>
          <b-tab title="Diff">
            <b-img fluid-grow rounded :src="images.diff"></b-img>
          </b-tab>
        </b-tabs>
      </b-col>
    </b-row>
  </div>
</template>


<script>
import api from '../app/api'

export default {
  name: 'Main',
  data() {
    return {
      activeTab: 2,
      testId: '',
      testName: '???',
      containerId: '???',
      percentage: '???',
      images: {
        reference: null,
        candidate: null,
        diff: null
      }
    }
  },

  mounted() {
    if (this.testIdParam) {
      this.testId = this.testIdParam;
      this.update();
    }
  },

  computed: {
    testIdParam() {
      return this.$route.params?.testIdParam
    }
  },

  watch: {
    containerId(value) {
      this.$emit('change-container-id', value)
    }
  },

  methods: {
    update() {
      if (this.testId) {
        if (this.testId !== this.testIdParam) {
          this.$router.push({
                name: 'test',
                params: {testIdParam: this.testId},
              }
          )

          this.getData(this.testId);
        } else {
          this.getData(this.testId);
        }
      }
    },
    getData(testId) {
      api.getTest(testId)
          .then(response => {
            const imagesList = response.data.images;
            for (let key in imagesList) {
              imagesList[key] = api.getImageURL(imagesList[key]);
            }

            this.images = imagesList;
            this.percentage = response.data.percentage;
            this.testName = response.data.test_container_name;
            this.containerId = response.data.test_container_id;
          })
          .catch(error => {
            // Restore default values on error
            this.images = {
              reference: null,
              candidate: null,
              diff: null
            };
            this.percentage = '???';
            this.testName = '???';
            this.containerId = '???';
          })
    }
  }
}

</script>
