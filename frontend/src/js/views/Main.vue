<template>
    <div>
        <b-row>
            <b-col>
                <b-form-input v-model="testId" @keyup.enter="update()" placeholder="Test Id"></b-form-input>
            </b-col>
            <b-col>
                <b-button block variant="info" @click="update">Update</b-button>
            </b-col>
        </b-row>

        <b-row>
            <b-col class="mt-3 mb-3">
                <h2>Results diff →
                    <percentage :value="percentage"></percentage>
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
import axios from 'axios';

const URL = 'http://127.0.0.1:3333/api/v1';
// const URL = '/test_data/';

export default {
    name: 'Main',
    data() {
        return {
            activeTab: 2,
            testId: '',
            percentage: '???',
            images: {
                reference: null,
                candidate: null,
                diff: null
            }
        }
    },
    props: ['testIdParam'],
    mounted() {
        if (this.testIdParam) {
            this.testId = this.testIdParam;
            this.update();
        }
    },
    methods: {
        update() {
            if (this.testId) {
                if (this.testId !== this.testIdParam) {
                    this.$router.push({
                        name: 'project',
                        params: {testIdParam: this.testId}
                    })

                    this.getData(this.testId);
                } else {
                    this.getData(this.testId);
                }
            }
        },
        getData(testId) {
            axios
                .get(URL + '/test/' + testId)
                .then(response => {
                    const imagesList = response.data.images;
                    for (let key in imagesList) {
                        imagesList[key] = URL + '/image/' + imagesList[key];
                    }

                    this.images = imagesList;
                    this.percentage = response.data.percentage;
                })
                .catch(error => {
                    // Restore default values on error
                    this.images = {
                        reference: null,
                        candidate: null,
                        diff: null
                    };
                    this.percentage = '???';
                })
        }
    }
}

</script>
