"use strict";
(self["webpackChunkimage_comparator_frontend"] = self["webpackChunkimage_comparator_frontend"] || []).push([["/js/app"],{

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-5[0].rules[0].use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./src/js/components/Percentage.vue?vue&type=script&lang=js&":
/*!***********************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-5[0].rules[0].use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./src/js/components/Percentage.vue?vue&type=script&lang=js& ***!
  \***********************************************************************************************************************************************************************************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   "default": () => (__WEBPACK_DEFAULT_EXPORT__)
/* harmony export */ });
//
//
//
//
//
//
//
/* harmony default export */ const __WEBPACK_DEFAULT_EXPORT__ = ({
  name: 'Percentage',
  props: ['value'],
  computed: {
    className: function className() {
      if (this.value === '???') {
        // @here you'll probably want to adjust color thresholds
        return 'percentage-unknown';
      } else if (this.value >= 99) {
        return 'percentage-awesome';
      } else if (this.value >= 95) {
        return 'percentage-good';
      } else if (this.value >= 70) {
        return 'percentage-bad';
      }

      return 'percentage-awful';
    },
    postfix: function postfix() {
      if (this.value === '???') {
        return '';
      }

      if (this.value === 100) {
        return '% 🎉🎉🎉';
      }

      return '%';
    }
  }
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-5[0].rules[0].use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./src/js/views/Main.vue?vue&type=script&lang=js&":
/*!************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-5[0].rules[0].use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./src/js/views/Main.vue?vue&type=script&lang=js& ***!
  \************************************************************************************************************************************************************************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   "default": () => (__WEBPACK_DEFAULT_EXPORT__)
/* harmony export */ });
/* harmony import */ var axios__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! axios */ "./node_modules/axios/index.js");
/* harmony import */ var axios__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(axios__WEBPACK_IMPORTED_MODULE_0__);
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//

var URL = 'http://127.0.0.1:3333/api/v1'; // @here change local URL to remote URL

/* harmony default export */ const __WEBPACK_DEFAULT_EXPORT__ = ({
  name: 'Main',
  data: function data() {
    return {
      projectId: '',
      // @here you can set default value, if you want
      testId: '',
      percentage: '???',
      images: {
        reference: null,
        candidate: null,
        diff: null
      }
    };
  },
  props: ['projectIdParam', 'testIdParam'],
  mounted: function mounted() {
    if (this.projectIdParam && this.testIdParam) {
      this.projectId = this.projectIdParam;
      this.testId = this.testIdParam;
      this.update();
    }
  },
  components: {},
  methods: {
    update: function update() {
      if (this.projectId && this.testId) {
        if (this.projectId !== this.projectIdParam || this.testId !== this.testIdParam) {
          this.$router.push({
            name: 'project',
            params: {
              projectIdParam: this.projectId,
              testIdParam: this.testId
            }
          });
          this.getData(this.projectId, this.testId);
        } else {
          this.getData(this.projectId, this.testId);
        }
      }
    },
    getData: function getData(projectId, testId) {
      var _this = this;

      axios__WEBPACK_IMPORTED_MODULE_0___default().get(URL + "/test/" + testId) // @here you can change URL params format
      .then(function (response) {
        var imagesList = response.data.images;

        for (var key in imagesList) {
          imagesList[key] = URL + '/image/' + imagesList[key];
        }

        _this.images = imagesList;
        _this.percentage = response.data.percentage;
      });
    }
  }
});

/***/ }),

/***/ "./src/js/app.js":
/*!***********************!*\
  !*** ./src/js/app.js ***!
  \***********************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var vue__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! vue */ "./node_modules/vue/dist/vue.esm.js");
/* harmony import */ var _app_ui__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./app/ui */ "./src/js/app/ui.js");
/* harmony import */ var _app_router__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./app/router */ "./src/js/app/router.js");
/* harmony import */ var _views_App__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./views/App */ "./src/js/views/App.vue");




/** Global Components */

_app_ui__WEBPACK_IMPORTED_MODULE_0__["default"].register();
/** Application */

new vue__WEBPACK_IMPORTED_MODULE_3__["default"]({
  router: _app_router__WEBPACK_IMPORTED_MODULE_1__["default"],
  render: function render(h) {
    return h(_views_App__WEBPACK_IMPORTED_MODULE_2__["default"]);
  }
}).$mount('#jameson-app');

/***/ }),

/***/ "./src/js/app/router.js":
/*!******************************!*\
  !*** ./src/js/app/router.js ***!
  \******************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   "default": () => (__WEBPACK_DEFAULT_EXPORT__)
/* harmony export */ });
/* harmony import */ var vue__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! vue */ "./node_modules/vue/dist/vue.esm.js");
/* harmony import */ var vue_router__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! vue-router */ "./node_modules/vue-router/dist/vue-router.esm.js");
/* harmony import */ var _views_Main__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ../views/Main */ "./src/js/views/Main.vue");
/* harmony import */ var _views_NotFound__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ../views/NotFound */ "./src/js/views/NotFound.vue");




vue__WEBPACK_IMPORTED_MODULE_2__["default"].use(vue_router__WEBPACK_IMPORTED_MODULE_3__["default"]);
var routes = [{
  path: '/',
  name: 'default',
  component: _views_Main__WEBPACK_IMPORTED_MODULE_0__["default"]
}, {
  path: '/project/:projectIdParam/test/:testIdParam',
  // http://my-service.com/project/my-project/test/test-id
  name: 'project',
  component: _views_Main__WEBPACK_IMPORTED_MODULE_0__["default"],
  props: true
}, {
  path: '*',
  component: _views_NotFound__WEBPACK_IMPORTED_MODULE_1__["default"]
}];
var router = new vue_router__WEBPACK_IMPORTED_MODULE_3__["default"]({
  // mode: 'history',
  scrollBehavior: function scrollBehavior() {
    return {
      x: 0,
      y: 0
    };
  },
  routes: routes
});
/* harmony default export */ const __WEBPACK_DEFAULT_EXPORT__ = (router);

/***/ }),

/***/ "./src/js/app/ui.js":
/*!**************************!*\
  !*** ./src/js/app/ui.js ***!
  \**************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   "default": () => (__WEBPACK_DEFAULT_EXPORT__)
/* harmony export */ });
/* harmony import */ var vue__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! vue */ "./node_modules/vue/dist/vue.esm.js");
/* harmony import */ var bootstrap_vue__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! bootstrap-vue */ "./node_modules/bootstrap-vue/esm/index.js");
/* harmony import */ var _components_Percentage__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ../components/Percentage */ "./src/js/components/Percentage.vue");



/** Register Global Components */

/* harmony default export */ const __WEBPACK_DEFAULT_EXPORT__ = ({
  register: function register() {
    vue__WEBPACK_IMPORTED_MODULE_1__["default"].use(bootstrap_vue__WEBPACK_IMPORTED_MODULE_2__.BootstrapVue);
    vue__WEBPACK_IMPORTED_MODULE_1__["default"].component('percentage', _components_Percentage__WEBPACK_IMPORTED_MODULE_0__["default"]);
  }
});

/***/ }),

/***/ "./src/sass/app.scss":
/*!***************************!*\
  !*** ./src/sass/app.scss ***!
  \***************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./src/js/components/Percentage.vue":
/*!******************************************!*\
  !*** ./src/js/components/Percentage.vue ***!
  \******************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   "default": () => (__WEBPACK_DEFAULT_EXPORT__)
/* harmony export */ });
/* harmony import */ var _Percentage_vue_vue_type_template_id_3ed17f66___WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./Percentage.vue?vue&type=template&id=3ed17f66& */ "./src/js/components/Percentage.vue?vue&type=template&id=3ed17f66&");
/* harmony import */ var _Percentage_vue_vue_type_script_lang_js___WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./Percentage.vue?vue&type=script&lang=js& */ "./src/js/components/Percentage.vue?vue&type=script&lang=js&");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! !../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");





/* normalize component */
;
var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_2__["default"])(
  _Percentage_vue_vue_type_script_lang_js___WEBPACK_IMPORTED_MODULE_1__["default"],
  _Percentage_vue_vue_type_template_id_3ed17f66___WEBPACK_IMPORTED_MODULE_0__.render,
  _Percentage_vue_vue_type_template_id_3ed17f66___WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  null,
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "src/js/components/Percentage.vue"
/* harmony default export */ const __WEBPACK_DEFAULT_EXPORT__ = (component.exports);

/***/ }),

/***/ "./src/js/views/App.vue":
/*!******************************!*\
  !*** ./src/js/views/App.vue ***!
  \******************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   "default": () => (__WEBPACK_DEFAULT_EXPORT__)
/* harmony export */ });
/* harmony import */ var _App_vue_vue_type_template_id_e5571e1a___WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./App.vue?vue&type=template&id=e5571e1a& */ "./src/js/views/App.vue?vue&type=template&id=e5571e1a&");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! !../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");

var script = {}


/* normalize component */
;
var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_1__["default"])(
  script,
  _App_vue_vue_type_template_id_e5571e1a___WEBPACK_IMPORTED_MODULE_0__.render,
  _App_vue_vue_type_template_id_e5571e1a___WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  null,
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "src/js/views/App.vue"
/* harmony default export */ const __WEBPACK_DEFAULT_EXPORT__ = (component.exports);

/***/ }),

/***/ "./src/js/views/Main.vue":
/*!*******************************!*\
  !*** ./src/js/views/Main.vue ***!
  \*******************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   "default": () => (__WEBPACK_DEFAULT_EXPORT__)
/* harmony export */ });
/* harmony import */ var _Main_vue_vue_type_template_id_4c84e352___WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./Main.vue?vue&type=template&id=4c84e352& */ "./src/js/views/Main.vue?vue&type=template&id=4c84e352&");
/* harmony import */ var _Main_vue_vue_type_script_lang_js___WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./Main.vue?vue&type=script&lang=js& */ "./src/js/views/Main.vue?vue&type=script&lang=js&");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! !../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");





/* normalize component */
;
var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_2__["default"])(
  _Main_vue_vue_type_script_lang_js___WEBPACK_IMPORTED_MODULE_1__["default"],
  _Main_vue_vue_type_template_id_4c84e352___WEBPACK_IMPORTED_MODULE_0__.render,
  _Main_vue_vue_type_template_id_4c84e352___WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  null,
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "src/js/views/Main.vue"
/* harmony default export */ const __WEBPACK_DEFAULT_EXPORT__ = (component.exports);

/***/ }),

/***/ "./src/js/views/NotFound.vue":
/*!***********************************!*\
  !*** ./src/js/views/NotFound.vue ***!
  \***********************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   "default": () => (__WEBPACK_DEFAULT_EXPORT__)
/* harmony export */ });
/* harmony import */ var _NotFound_vue_vue_type_template_id_2917530d___WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./NotFound.vue?vue&type=template&id=2917530d& */ "./src/js/views/NotFound.vue?vue&type=template&id=2917530d&");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! !../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");

var script = {}


/* normalize component */
;
var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_1__["default"])(
  script,
  _NotFound_vue_vue_type_template_id_2917530d___WEBPACK_IMPORTED_MODULE_0__.render,
  _NotFound_vue_vue_type_template_id_2917530d___WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  null,
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "src/js/views/NotFound.vue"
/* harmony default export */ const __WEBPACK_DEFAULT_EXPORT__ = (component.exports);

/***/ }),

/***/ "./src/js/components/Percentage.vue?vue&type=script&lang=js&":
/*!*******************************************************************!*\
  !*** ./src/js/components/Percentage.vue?vue&type=script&lang=js& ***!
  \*******************************************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   "default": () => (__WEBPACK_DEFAULT_EXPORT__)
/* harmony export */ });
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_5_0_rules_0_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_Percentage_vue_vue_type_script_lang_js___WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-5[0].rules[0].use[0]!../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./Percentage.vue?vue&type=script&lang=js& */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-5[0].rules[0].use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./src/js/components/Percentage.vue?vue&type=script&lang=js&");
 /* harmony default export */ const __WEBPACK_DEFAULT_EXPORT__ = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_5_0_rules_0_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_Percentage_vue_vue_type_script_lang_js___WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./src/js/views/Main.vue?vue&type=script&lang=js&":
/*!********************************************************!*\
  !*** ./src/js/views/Main.vue?vue&type=script&lang=js& ***!
  \********************************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   "default": () => (__WEBPACK_DEFAULT_EXPORT__)
/* harmony export */ });
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_5_0_rules_0_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_Main_vue_vue_type_script_lang_js___WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-5[0].rules[0].use[0]!../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./Main.vue?vue&type=script&lang=js& */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-5[0].rules[0].use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./src/js/views/Main.vue?vue&type=script&lang=js&");
 /* harmony default export */ const __WEBPACK_DEFAULT_EXPORT__ = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_5_0_rules_0_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_Main_vue_vue_type_script_lang_js___WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./src/js/components/Percentage.vue?vue&type=template&id=3ed17f66&":
/*!*************************************************************************!*\
  !*** ./src/js/components/Percentage.vue?vue&type=template&id=3ed17f66& ***!
  \*************************************************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   "render": () => (/* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Percentage_vue_vue_type_template_id_3ed17f66___WEBPACK_IMPORTED_MODULE_0__.render),
/* harmony export */   "staticRenderFns": () => (/* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Percentage_vue_vue_type_template_id_3ed17f66___WEBPACK_IMPORTED_MODULE_0__.staticRenderFns)
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Percentage_vue_vue_type_template_id_3ed17f66___WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./Percentage.vue?vue&type=template&id=3ed17f66& */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./src/js/components/Percentage.vue?vue&type=template&id=3ed17f66&");


/***/ }),

/***/ "./src/js/views/App.vue?vue&type=template&id=e5571e1a&":
/*!*************************************************************!*\
  !*** ./src/js/views/App.vue?vue&type=template&id=e5571e1a& ***!
  \*************************************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   "render": () => (/* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_App_vue_vue_type_template_id_e5571e1a___WEBPACK_IMPORTED_MODULE_0__.render),
/* harmony export */   "staticRenderFns": () => (/* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_App_vue_vue_type_template_id_e5571e1a___WEBPACK_IMPORTED_MODULE_0__.staticRenderFns)
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_App_vue_vue_type_template_id_e5571e1a___WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./App.vue?vue&type=template&id=e5571e1a& */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./src/js/views/App.vue?vue&type=template&id=e5571e1a&");


/***/ }),

/***/ "./src/js/views/Main.vue?vue&type=template&id=4c84e352&":
/*!**************************************************************!*\
  !*** ./src/js/views/Main.vue?vue&type=template&id=4c84e352& ***!
  \**************************************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   "render": () => (/* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Main_vue_vue_type_template_id_4c84e352___WEBPACK_IMPORTED_MODULE_0__.render),
/* harmony export */   "staticRenderFns": () => (/* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Main_vue_vue_type_template_id_4c84e352___WEBPACK_IMPORTED_MODULE_0__.staticRenderFns)
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Main_vue_vue_type_template_id_4c84e352___WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./Main.vue?vue&type=template&id=4c84e352& */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./src/js/views/Main.vue?vue&type=template&id=4c84e352&");


/***/ }),

/***/ "./src/js/views/NotFound.vue?vue&type=template&id=2917530d&":
/*!******************************************************************!*\
  !*** ./src/js/views/NotFound.vue?vue&type=template&id=2917530d& ***!
  \******************************************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   "render": () => (/* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_NotFound_vue_vue_type_template_id_2917530d___WEBPACK_IMPORTED_MODULE_0__.render),
/* harmony export */   "staticRenderFns": () => (/* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_NotFound_vue_vue_type_template_id_2917530d___WEBPACK_IMPORTED_MODULE_0__.staticRenderFns)
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_NotFound_vue_vue_type_template_id_2917530d___WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./NotFound.vue?vue&type=template&id=2917530d& */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./src/js/views/NotFound.vue?vue&type=template&id=2917530d&");


/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./src/js/components/Percentage.vue?vue&type=template&id=3ed17f66&":
/*!****************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./src/js/components/Percentage.vue?vue&type=template&id=3ed17f66& ***!
  \****************************************************************************************************************************************************************************************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   "render": () => (/* binding */ render),
/* harmony export */   "staticRenderFns": () => (/* binding */ staticRenderFns)
/* harmony export */ });
var render = function () {
  var _vm = this
  var _h = _vm.$createElement
  var _c = _vm._self._c || _h
  return _c("span", { class: _vm.className }, [
    _vm._v("\n    " + _vm._s(_vm.value) + " " + _vm._s(_vm.postfix) + "\n"),
  ])
}
var staticRenderFns = []
render._withStripped = true



/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./src/js/views/App.vue?vue&type=template&id=e5571e1a&":
/*!****************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./src/js/views/App.vue?vue&type=template&id=e5571e1a& ***!
  \****************************************************************************************************************************************************************************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   "render": () => (/* binding */ render),
/* harmony export */   "staticRenderFns": () => (/* binding */ staticRenderFns)
/* harmony export */ });
var render = function () {
  var _vm = this
  var _h = _vm.$createElement
  var _c = _vm._self._c || _h
  return _c(
    "b-container",
    { staticClass: "bv-example-row" },
    [
      _c(
        "b-row",
        [
          _c("b-col", { staticClass: "mt-3 mb-3" }, [
            _c("h1", [_vm._v("Jameson the Image Comparator")]),
          ]),
        ],
        1
      ),
      _vm._v(" "),
      _c("router-view"),
    ],
    1
  )
}
var staticRenderFns = []
render._withStripped = true



/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./src/js/views/Main.vue?vue&type=template&id=4c84e352&":
/*!*****************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./src/js/views/Main.vue?vue&type=template&id=4c84e352& ***!
  \*****************************************************************************************************************************************************************************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   "render": () => (/* binding */ render),
/* harmony export */   "staticRenderFns": () => (/* binding */ staticRenderFns)
/* harmony export */ });
var render = function () {
  var _vm = this
  var _h = _vm.$createElement
  var _c = _vm._self._c || _h
  return _c(
    "div",
    [
      _c(
        "b-row",
        [
          _c(
            "b-col",
            [
              _c("b-form-input", {
                attrs: { placeholder: "Project ID" },
                model: {
                  value: _vm.projectId,
                  callback: function ($$v) {
                    _vm.projectId = $$v
                  },
                  expression: "projectId",
                },
              }),
            ],
            1
          ),
          _vm._v(" "),
          _c(
            "b-col",
            [
              _c("b-form-input", {
                attrs: { placeholder: "Test Id" },
                model: {
                  value: _vm.testId,
                  callback: function ($$v) {
                    _vm.testId = $$v
                  },
                  expression: "testId",
                },
              }),
            ],
            1
          ),
          _vm._v(" "),
          _c(
            "b-col",
            [
              _c(
                "b-button",
                {
                  attrs: { block: "", variant: "info" },
                  on: { click: _vm.update },
                },
                [_vm._v("Update")]
              ),
            ],
            1
          ),
        ],
        1
      ),
      _vm._v(" "),
      _c(
        "b-row",
        [
          _c("b-col", { staticClass: "mt-3 mb-3" }, [
            _c(
              "h2",
              [
                _vm._v("Results → "),
                _c("percentage", { attrs: { value: _vm.percentage } }),
              ],
              1
            ),
          ]),
        ],
        1
      ),
      _vm._v(" "),
      _c(
        "b-row",
        [
          _c(
            "b-col",
            [
              _c("h3", [_vm._v("Reference")]),
              _vm._v(" "),
              _c("b-img", {
                attrs: {
                  "fluid-grow": "",
                  rounded: "",
                  src: _vm.images.reference,
                },
              }),
            ],
            1
          ),
          _vm._v(" "),
          _c(
            "b-col",
            [
              _c("h3", [_vm._v("Diff ")]),
              _vm._v(" "),
              _c("b-img", {
                attrs: { "fluid-grow": "", rounded: "", src: _vm.images.diff },
              }),
            ],
            1
          ),
          _vm._v(" "),
          _c(
            "b-col",
            [
              _c("h3", [_vm._v("Candidate")]),
              _vm._v(" "),
              _c("b-img", {
                attrs: {
                  "fluid-grow": "",
                  rounded: "",
                  src: _vm.images.candidate,
                },
              }),
            ],
            1
          ),
        ],
        1
      ),
    ],
    1
  )
}
var staticRenderFns = []
render._withStripped = true



/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./src/js/views/NotFound.vue?vue&type=template&id=2917530d&":
/*!*********************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./src/js/views/NotFound.vue?vue&type=template&id=2917530d& ***!
  \*********************************************************************************************************************************************************************************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   "render": () => (/* binding */ render),
/* harmony export */   "staticRenderFns": () => (/* binding */ staticRenderFns)
/* harmony export */ });
var render = function () {
  var _vm = this
  var _h = _vm.$createElement
  var _c = _vm._self._c || _h
  return _c(
    "b-row",
    [_c("b-col", [_c("h2", [_vm._v("Oops. Page not found.")])])],
    1
  )
}
var staticRenderFns = []
render._withStripped = true



/***/ })

},
/******/ __webpack_require__ => { // webpackRuntimeModules
/******/ var __webpack_exec__ = (moduleId) => (__webpack_require__(__webpack_require__.s = moduleId))
/******/ __webpack_require__.O(0, ["css/app","/js/vendor"], () => (__webpack_exec__("./src/js/app.js"), __webpack_exec__("./src/sass/app.scss")));
/******/ var __webpack_exports__ = __webpack_require__.O();
/******/ }
]);