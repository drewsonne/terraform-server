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
    // AMD.
    define(['expect.js', '../../src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require('../../src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.Lunaform);
  }
}(this, function(expect, Lunaform) {
  'use strict';

  var instance;

  beforeEach(function() {
    instance = new Lunaform.ResourceAuthUser();
  });

  var getProperty = function(object, getter, property) {
    // Use getter method if present; otherwise, get the property directly.
    if (typeof object[getter] === 'function')
      return object[getter]();
    else
      return object[property];
  }

  var setProperty = function(object, setter, property, value) {
    // Use setter method if present; otherwise, set the property directly.
    if (typeof object[setter] === 'function')
      object[setter](value);
    else
      object[property] = value;
  }

  describe('ResourceAuthUser', function() {
    it('should create an instance of ResourceAuthUser', function() {
      // uncomment below and update the code to test ResourceAuthUser
      //var instane = new Lunaform.ResourceAuthUser();
      //expect(instance).to.be.a(Lunaform.ResourceAuthUser);
    });

    it('should have the property shortname (base name: "shortname")', function() {
      // uncomment below and update the code to test the property shortname
      //var instane = new Lunaform.ResourceAuthUser();
      //expect(instance).to.be();
    });

    it('should have the property name (base name: "name")', function() {
      // uncomment below and update the code to test the property name
      //var instane = new Lunaform.ResourceAuthUser();
      //expect(instance).to.be();
    });

    it('should have the property id (base name: "id")', function() {
      // uncomment below and update the code to test the property id
      //var instane = new Lunaform.ResourceAuthUser();
      //expect(instance).to.be();
    });

    it('should have the property groups (base name: "groups")', function() {
      // uncomment below and update the code to test the property groups
      //var instane = new Lunaform.ResourceAuthUser();
      //expect(instance).to.be();
    });

    it('should have the property apiKeys (base name: "api-keys")', function() {
      // uncomment below and update the code to test the property apiKeys
      //var instane = new Lunaform.ResourceAuthUser();
      //expect(instance).to.be();
    });

  });

}));
