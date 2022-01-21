<template>
    <div>
        <b-row>
            <b-col>
                <b-form-input v-model="projectId" placeholder="Project ID"></b-form-input>
            </b-col>
            <b-col>
                <b-form-input v-model="testId" placeholder="Test Id"></b-form-input>
            </b-col>
            <b-col>
                <b-button block variant="info" @click="update">Update</b-button>
            </b-col>
        </b-row>

        <b-row>
            <b-col class="mt-3 mb-3">
                <h2>Results → <percentage :value="percentage"></percentage></h2>
            </b-col>
        </b-row>

        <b-row>
            <b-col>
                <h3>Reference</h3>
                <b-img fluid-grow rounded :src="images.reference"></b-img>
            </b-col>
            <b-col>
                <h3>Diff </h3>
                <b-img fluid-grow rounded :src="images.diff"></b-img>
            </b-col>
            <b-col>
                <h3>Candidate</h3>
                <b-img fluid-grow rounded :src="images.candidate"></b-img>
            </b-col>
        </b-row>
    </div>
</template>


<script>
import axios from 'axios';

const URL = '/test_data/response.json'; // @here change local URL to remote URL

export default {
    name: 'Main',
    data() {
        return {
            projectId: '', // @here you can set default value, if you want
            testId: '',
            percentage: '???',
            images: {
                reference: null,
                candidate: null,
                diff: null
            }
        }
    },
    props: ['projectIdParam', 'testIdParam'],
    mounted() {
        if (this.projectIdParam && this.testIdParam) {
            this.projectId = this.projectIdParam;
            this.testId = this.testIdParam;

            this.update();
        }
    },
    components: {},
    methods: {
        update() {
            if (this.projectId && this.testId) {
                if (this.projectId !== this.projectIdParam || this.testId !== this.testIdParam) {
                    this.$router.push({
                        name: 'project',
                        params: {projectIdParam: this.projectId, testIdParam: this.testId}
                    })

                    this.getData(this.projectId, this.testId);
                } else {
                    this.getData(this.projectId, this.testId);
                }
            }
        },
        getData(projectId, testId) {
            axios
                .get(URL + "?projectId=" + projectId + "&testId=" + testId) // @here you can change URL params format
                .then(response => {
                    this.images = response.data.images;
                    this.percentage = response.data.percentage;
                })
                .catch(function (error) {
                });
        }
    }
}

</script>
