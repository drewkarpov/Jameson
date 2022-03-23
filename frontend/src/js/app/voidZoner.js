import api from './api';

const FILL_COLOR = 'cyan';
const STROKE_COLOR = 'magenta';

export default class VoidZoner {
  refs;
  rects;
  containerData;
  scaleProportion;

  constructor(responseData, refs) {
    this.containerData = responseData;
    this.refs = refs;
  }

  loadImage(callback) {
    const image = new window.Image();
    image.src = api.getImageURL(this.containerData.reference_id);
    image.onload = () => {
      let canvasRect = this.refs.canvasContainer?.getBoundingClientRect();
      let stageConfig, rects;

      if (canvasRect) {
        this.scaleProportion = canvasRect.width / image.width;

        image.height = (image.height * canvasRect.width) / image.width;
        image.width = canvasRect.width;

        stageConfig = {
          width: canvasRect.width,
          height: image.height
        };

        if (this.containerData.void_zones && this.containerData.void_zones.length > 0) {
          rects = this.renderRects(this.containerData.void_zones);
        }
      }

      callback(image, stageConfig, rects);
    };
  }

  renderRects(list) {
    return list.map((rect) => {
      let defaultRectConfig = {
        x: 0,
        y: 0,
        width: 0,
        height: 0,
        fill: FILL_COLOR,
        stroke: STROKE_COLOR,
        strokeWidth: 3,
      };

      defaultRectConfig.x = Math.ceil(rect.offset_x * this.scaleProportion);
      defaultRectConfig.y = Math.ceil(rect.offset_y * this.scaleProportion);
      defaultRectConfig.width = Math.ceil(rect.width * this.scaleProportion);
      defaultRectConfig.height = Math.ceil(rect.height * this.scaleProportion);

      return defaultRectConfig;
    });
  }

};