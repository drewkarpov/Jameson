import api from './api';

const FILL_COLOR = 'transparent';
const STROKE_COLOR = 'blue';

export default class VoidZoner {
  width;
  height;

  containerData;
  canvasContainer;
  scaleFactor;

  isImageLoaded = false;
  isDrawing = false;
  isDrawingEnabled = false;

  constructor(responseData, refs) {
    this.containerData = responseData;
    this.canvasContainer = refs.canvasContainer;
  }

  loadImage(callback) {
    const image = new window.Image();
    image.src = api.getImageURL(this.containerData.reference_id);
    image.onload = () => {
      let canvasRect = this.canvasContainer?.getBoundingClientRect();
      let stageConfig, rects;

      if (canvasRect) {
        this.isImageLoaded = true;
        this.width = canvasRect.width;
        this.height = canvasRect.height;

        this.scaleFactor = this.width / image.width;

        image.height = (image.height * this.width) / image.width;
        image.width = this.width;

        stageConfig = {
          width: this.width,
          height: image.height
        };

        if (this.containerData.void_zones && this.containerData.void_zones.length > 0) {
          rects = this.renderRects(this.containerData.void_zones);
        }
      }

      callback(image, stageConfig, rects);
    };
  }

  renderRect(x, y, width, height) {
    return {
      x: x,
      y: y,
      width: width,
      height: height,
      fill: FILL_COLOR,
      stroke: STROKE_COLOR,
      strokeWidth: 3,
      draggable: true
    };
  }

  renderRects(list) {
    return list.map((rect) => {
      return this.renderRect(
          Math.ceil(rect.offset_x * this.scaleFactor),
          Math.ceil(rect.offset_y * this.scaleFactor),
          Math.ceil(rect.width * this.scaleFactor),
          Math.ceil(rect.height * this.scaleFactor)
      );
    });
  }

  prepareRects(list) {
    return list
        .filter((rect) => (rect.width > 2 && rect.height > 2)) // Filter missclicks
        .map((rect) => {
          let rectData = {
            offset_x: 0,
            offset_y: 0,
            width: 0,
            height: 0
          };

          if (this.scaleFactor !== 0) {
            rectData.offset_x = Math.round(rect.x / this.scaleFactor);
            rectData.offset_y = Math.round(rect.y / this.scaleFactor);
            rectData.width = Math.round(rect.width / this.scaleFactor);
            rectData.height = Math.round(rect.height / this.scaleFactor);
          }

          return rectData;
        });
  }

};