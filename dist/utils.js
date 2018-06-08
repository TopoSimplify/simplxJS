

/**
 * @description constructs a GeometryFactory that
 *  generates Geometries having a floating
 * @constructor
 */
function GeometryFactory() {
  "use strict";
  if(!(this instanceof GeometryFactory)){
    return new GeometryFactory();
  }
}

/**
 * Creates a Point using the given Coordinate; a null Coordinate will create an
 * empty Geometry.
 * @return {Array} A new Point.
 */
GeometryFactory.prototype.createPoint = function () {
  return Array.prototype.slice.call(arguments);
};

/**
 * @description creates a LineString using the given Coordinates; a null or empty array will
 * create an empty LineString. Consecutive points must not be equal.
 * @param coordinates{[]} - coordinates an array of coordinates
 * @return {Array} A new LineString.
 */
GeometryFactory.prototype.createLineString = function (coordinates) {
  return coordinates;
};

/**
 * @description creates a LinearRing using the given Coordinates; a null or empty array will
 * create an empty LinearRing. The points must form a closed and simple
 * linestring. Consecutive points must not be equal.
 * @param coordinates{[]} -  array of coordinates
 * @return {Array} A new LinearRing.
 */
GeometryFactory.prototype.createLinearRing = function (coordinates) {
  return  coordinates;
};

/**
 * @description constructs a <code>Polygon</code> with the given exterior boundary and
 * interior boundaries.
 * @param shell{Array} - an array of coordinates
 *          shell will be used as an outer boundary of the new <code>Polygon</code>
 * @param [holes]{Array[]} - an array of arrays of shells forming hole boundry
 *          holes the inner boundaries of the new <code>Polygon</code>
 * @return {Array} A new Polygon.
 */
GeometryFactory.prototype.createPolygon = function (shell, holes) {
  return [shell, holes];
};
