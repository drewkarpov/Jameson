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
            <b-col class="mt-3">
                <h2 class="mb-3" v-if="loaded">
                    {{ container.name }}
                    <b-badge variant="light" pill v-if="container.approved">Approved</b-badge>
                </h2>

                <b-button-group v-if="loaded">
                    <b-button :variant="rectangleButtonVariant" @click="rectangle">
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
                        ref="stage"
                        :config="stageConfig"
                        @mousemove="handleMouseMove"
                        @mouseDown="handleMouseDown"
                        @mouseUp="handleMouseUp">
                        <v-layer>
                            <v-image :config="{image: image}"/>

                            <v-rect
                                v-for="(rect, index) in rects"
                                :key="index"
                                :config="rect"
                                @dragend="handleDragEnd"
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
            voidZoner: null,
            container: null,
            containerId: null,
            image: null,
            stageConfig: {},
            rects: []
        }
    },
    computed: {
        rectangleButtonVariant: function () {
            if (this.voidZoner?.isDrawingEnabled) {
                return 'primary';
            }

            return 'outline-primary';
        },

        containerIdParam() {
            return this.$route.params?.containerIdParam
        }
    },
    mounted() {
        if (this.containerIdParam) {
            this.containerId = this.containerIdParam;
            this.update();
        }
    },

    methods: {
        rectangle() {
            if (this.voidZoner) {
                this.voidZoner.isDrawingEnabled = true;
            }
        },
        handleMouseDown() {
            if (this.voidZoner?.isDrawingEnabled) {
                this.voidZoner.isDrawing = true;

                const pointerPosition = this.$refs.stage.getNode().getPointerPosition();
                this.rects = [
                    ...this.rects,
                    this.voidZoner.renderRect(pointerPosition.x, pointerPosition.y, 0, 0)
                ];
            }
        },
        handleMouseUp() {
            if (this.voidZoner?.isDrawing) {
                this.saveData();
            }

            this.voidZoner.isDrawing = false;
            this.voidZoner.isDrawingEnabled = false;
        },
        handleMouseMove() {
            if (!this.voidZoner?.isDrawingEnabled || !this.voidZoner?.isDrawing) {
                return;
            }

            const pointerPosition = this.$refs.stage.getNode().getPointerPosition();

            let latestRectangle = this.rects[this.rects.length - 1];
            if (latestRectangle) {
                latestRectangle.width = pointerPosition.x - latestRectangle.x;
                latestRectangle.height = pointerPosition.y - latestRectangle.y;
            }
        },
        handleDragEnd(event) {
            // Update rect position in array
            let targetIndex = event.target?.index;
            if (targetIndex && this.rects[targetIndex - 1]) {
                let newX = event.target?.attrs?.x;
                let newY = event.target?.attrs?.y;

                this.rects[targetIndex - 1].x = newX;
                this.rects[targetIndex - 1].y = newY;
            }

            this.saveData();
        },
        clear() {
            this.rects = [];
            this.saveData();
        },
        update() {
            this.cleanUp();

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
        cleanUp() {
            this.image = null;
            this.rects = [];
            this.loaded = false;
        },
        approve() {
            api.approveContainer(this.containerId).then((data) => {
                this.container.approved = true;
                this.notify('Okay', 'Container Approved');
            }).catch((e) => {
                this.notify('Error', 'Some error occurred', 'danger');
            })
        },
        getData(containerId) {
            api
                .getContainer(containerId)
                .then(response => {
                    this.container = response.data;
                    this.voidZoner = new VoidZoner(response.data, this.$refs);

                    this.voidZoner.loadImage((image, stage, rects) => {
                        this.image = image;
                        this.stageConfig = stage;
                        this.rects = rects || [];

                        this.loaded = true;
                    });
                })
                .catch(error => {
                    this.loaded = false;
                    this.container = null;
                    this.image = null;
                })
        },
        saveData() {
            let voidZones = this.voidZoner.prepareRects(this.rects);
            api.setContainerVoidZones(this.containerId, voidZones);
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
