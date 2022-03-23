<template>
    <div>
        <b-row>
            <b-col>
                <b-form-input v-model="containerId" @keyup.enter="update" placeholder="Container Id"></b-form-input>
            </b-col>
            <b-col>
                <b-button block variant="primary" @click="update">Update</b-button>
            </b-col>
        </b-row>

        <b-row>
            <b-col class="mt-3" v-if="loaded">
                <h2 class="mb-3">
                    {{ container.name }}
                    <b-badge variant="light" pill v-if="container.approved">Approved</b-badge>
                </h2>

                <b-button-group>
                    <b-button variant="outline-primary" @click="rectangle">
                        <b-icon icon="bounding-box-circles" aria-hidden="true"></b-icon>
                        Rectangle
                    </b-button>
                    <b-button variant="outline-danger" @click="clear">
                        <b-icon icon="trash" aria-hidden="true"></b-icon>
                        Clear all rectangles
                    </b-button>
                    <b-button variant="outline-success" @click="approve" v-if="!container.approved">
                        <b-icon icon="check2-circle" aria-hidden="true"></b-icon>
                        Approve container
                    </b-button>
                </b-button-group>

                <div class="mt-3" ref="canvasContainer">
                    <v-stage
                        :config="stageConfig"
                        @mousemove="handleMouseMove"
                        @mouseDown="handleMouseDown"
                        @mouseUp="handleMouseUp">
                        <v-layer>
                            <v-image :config="{image: image}"/>

                            <v-rect
                                v-for="(rec, index) in rects"
                                :key="index"
                                :config="rec"
                            />
                        </v-layer>
                    </v-stage>
                </div>
            </b-col>
        </b-row>
    </div>
</template>


<script>
import api from '../app/api'
import VoidZoner from '../app/voidZoner';

export default {
    name: 'VoidZones',
    data() {
        return {
            loaded: false,
            container: null,
            containerId: null,
            image: null,
            stageConfig: {},
            rects: []
        }
    },
    props: ['containerIdParam'],
    mounted() {
        if (this.containerIdParam) {
            this.containerId = this.containerIdParam;
            this.update();
        }
    },
    methods: {
        rectangle() {
            // todo
        },
        handleMouseMove() {
            // todo
        },
        handleMouseDown() {
            // todo
        },
        handleMouseUp() {
            // todo
        },
        clear() {
            // todo
        },
        approve() {
            api.approveContainer(this.containerId).then((data) => {
                this.container.approved = true;
                this.notify('Okay', 'Container Approved');
            }).catch((e) => {
                this.notify('Error', 'Some error occurred', 'danger');
            })
        },
        update() {
            if (this.containerId) {
                if (this.containerId !== this.containerIdParam) {
                    this.$router.push({
                        name: 'container',
                        params: {containerIdParam: this.containerId}
                    })
                }

                this.getData(this.containerId);
            }
        },
        getData(containerId) {
            api
                .getContainer(containerId)
                .then(response => {
                    this.container = response.data;
                    const voidZoner = new VoidZoner(response.data, this.$refs);

                    voidZoner.loadImage((image, stage, rects) => {
                        this.image = image;
                        this.stageConfig = stage;
                        this.rects = rects;
                    });

                    this.loaded = true;
                })
                .catch(error => {
                    this.loaded = false;
                    this.container = null;
                    this.image = null;
                })
        },
        notify(title, message, variant = 'success') {
            this.$bvToast.toast(message, {
                title: title,
                autoHideDelay: 2000,
                appendToast: true,
                variant: variant
            })
        }
    }
}

</script>
