/**
 * lunaform
 * This is a RESTful server for managing Terraform plan and apply jobs and the auditing of actions to approve those apply jobs. The inspiration for this project is the AWS CloudFormation API's. The intention is to implement a locking mechanism not only for the terraform state, but for the plan and apply of terraform modules. Once a `module` plan starts, it is instantiated as a `stack` within the nomencalture of `lunaform`. 
 *
 * OpenAPI spec version: 0.0.1-alpha
 * Contact: drew.sonne@gmail.com
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.3.1
 *
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(['ApiClient'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'));
  } else {
    // Browser globals (root is window)
    if (!root.Lunaform) {
      root.Lunaform = {};
    }
    root.Lunaform.HalHref = factory(root.Lunaform.ApiClient);
  }
}(this, function(ApiClient) {
  'use strict';




  /**
   * The HalHref model module.
   * @module model/HalHref
   * @version 0.0.1-alpha
   */

  /**
   * Constructs a new <code>HalHref</code>.
   * @alias module:model/HalHref
   * @class
   */
  var exports = function() {
    var _this = this;


  };

  /**
   * Constructs a <code>HalHref</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/HalHref} obj Optional instance to populate.
   * @return {module:model/HalHref} The populated <code>HalHref</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();

      if (data.hasOwnProperty('href')) {
        obj['href'] = ApiClient.convertToType(data['href'], 'String');
      }
    }
    return obj;
  }

  /**
   * @member {String} href
   */
  exports.prototype['href'] = undefined;



  return exports;
}));


