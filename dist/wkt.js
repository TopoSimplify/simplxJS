/**
 * Adapted from OpenLayers 2.11 implementation.
 */
/**
 * @description converts a geometry in Well-Known Text format to a {@link Geometry}.
 * <p>
 * <code>WKTReader</code> supports extracting <code>Geometry</code> objects
 * from either {@link Reader}s or {@link String}s. This allows it to function
 * as a parser to read <code>Geometry</code> objects from text blocks embedded
 * in other data formats (e.g. XML).
 * <P>
 * <p>
 * A <code>WKTReader</code> is parameterized by a <code>GeometryFactory</code>,
 * to allow it to create <code>Geometry</code> objects of the appropriate
 * implementation. In particular, the <code>GeometryFactory</code> determines
 * the <code>PrecisionModel</code> and <code>SRID</code> that is used.
 * <P>
 *
 * @constructor
 */
function WKTReader(geomFactory) {
	"use strict";
	if(!(this instanceof WKTReader)) {
		return new WKTReader(geomFactory);
	}

	this.geomFactory = geomFactory || new GeometryFactory();
	this.regExes     = {
		'typeStr': /^\s*(\w+)\s*\(\s*(.*)\s*\)\s*$/,
		'emptyTypeStr': /^\s*(\w+)\s*EMPTY\s*$/,
		'spaces': /\s+/,
		'parenComma': /\)\s*,\s*\(/,
		'doubleParenComma': /\)\s*\)\s*,\s*\(\s*\(/, // can't use {2} here
		'trimParens': /^\s*\(?(.*?)\)?\s*$/
	};
}



/**
 * @description deserialize a WKT string and return a geometry. Supports WKT for POINT,
 * MULTIPOINT, LINESTRING, LINEARRING, MULTILINESTRING, POLYGON, MULTIPOLYGON,
 * and GEOMETRYCOLLECTION.
 * @param wkt{String} - wkt A WKT string.
 * @return {Geometry} A geometry instance.
 */
WKTReader.prototype.read = function(wkt) {
	var geometry, type, str;
	wkt         = wkt.replace(/[\n\r]/g, ' ');
	var matches = this.regExes.typeStr.exec(wkt);
	if(wkt.search('EMPTY') !== -1) {
		matches    = this.regExes.emptyTypeStr.exec(wkt);
		matches[2] = undefined;
	}
	if(matches) {
		type = matches[1].toLowerCase();
		str  = matches[2];
		if(this[type]) {
			geometry = this[type](str);
		}
	}

	if(geometry === undefined) {
		throw new Error('Could not parse WKT ' + wkt);
	}
	return geometry;
};

/**
 * Return point geometry given a point WKT fragment.
 * @param str{String} - str A WKT fragment representing the point.
 * @return {Point} A point geometry.
 * @private
 */
WKTReader.prototype.point = function(str) {
	if(str === undefined) {
		throw new Error("wkt point not passed");
	}
	var coords = str.trim().split(this.regExes.spaces);
	return this.geomFactory.createPoint(coords.map(parseFloat));
};

/**
 * Return a linestring geometry given a linestring WKT fragment.
 * @param str{String} - A WKT fragment representing the linestring.
 * @return {Array} A linestring coodinate list.
 * @private
 */
WKTReader.prototype.stringcoords = function(str) {
	if(str === undefined) {
		throw new Error("coordinate list not passed");
	}
	var points     = str.trim().split(',');
	var n          = points.length, coords;
	var components = new Array(n);
	for(var i = 0; i < n; ++i) {
		coords        = points[i].trim().split(this.regExes.spaces);
		components[i] = coords.map(parseFloat);
	}
	return components;
};
/**
 * Return a linestring geometry given a linestring WKT fragment.
 * @param str{String} - A WKT fragment representing the linestring.
 * @return {LineString} A linestring geometry.
 * @private
 */
WKTReader.prototype.linestring = function(str) {
	if(str === undefined) {
		throw new Error("coordinate list not passed");
	}
	return this.geomFactory.createLineString(
		this.stringcoords(str)
	);
};

/**
 * Return a linearring geometry given a linearring WKT fragment.
 * @param str{String} - A WKT fragment representing the linearring.
 * @return {LinearRing} A linearring geometry.
 * @private
 */
WKTReader.prototype.linearring = function(str) {
	if(str === undefined) {
		throw new Error("coordinate list is empty");
	}
	return this.geomFactory.createLinearRing(
		this.stringcoords(str)
	);
};

/**
 * Return a polygon geometry given a polygon WKT fragment.
 * @param str{String} - A WKT fragment representing the polygon.
 * @return {Array} A polygon geometry.
 * @private
 */
WKTReader.prototype.polygon = function(str) {
	if(str === undefined) {
		throw new Error('wkt fragement not passed')
	}
	var ring, components;
	var rings        = str.trim().split(this.regExes.parenComma);
	var n            = rings.length;
	var shell, holes = new Array(n - 1);
	for(var i = 0; i < n; i++) {
		ring       = rings[i].replace(this.regExes.trimParens, '$1');
		components = this.stringcoords(ring);
		(i === 0) && (shell = components);
		(i > 0) && (holes[i - 1] = components);
	}
	return this.geomFactory.createPolygon(shell, holes);
};

