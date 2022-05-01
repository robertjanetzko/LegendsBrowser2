var sitesLayer = L.layerGroup();
var constructionsLayer = L.layerGroup();
var landmassesLayer = L.layerGroup();
var regionsLayer = L.layerGroup();
var mountainsLayer = L.layerGroup();
var evilnessLayer = L.layerGroup();

var map = L.map('map', {
    maxZoom: 6,
    minZoom: 0,
    crs: L.CRS.Simple,
    layers: [sitesLayer, constructionsLayer, mountainsLayer, evilnessLayer]
});

map.getBoundsZoom = function (bounds, inside, padding) { // (LatLngBounds[, Boolean, Point]) -> Number
    bounds = L.latLngBounds(bounds);

    var zoom = this.getMinZoom() - (inside ? 1 : 0),
        maxZoom = this.getMaxZoom(),
        size = this.getSize(),

        nw = bounds.getNorthWest(),
        se = bounds.getSouthEast(),

        zoomNotFound = true,
        boundsSize;

    padding = L.point(padding || [0, 0]);

    var incement = 0.02;
    do {
        zoom += incement;
        boundsSize = this.project(se, zoom).subtract(this.project(nw, zoom)).add(padding).floor();
        zoomNotFound = !inside ? size.contains(boundsSize) : boundsSize.x < size.x || boundsSize.y < size.y;

    } while (zoomNotFound && zoom <= maxZoom);

    if (zoomNotFound && inside) {
        return null;
    }

    return inside ? zoom : zoom - incement;
}

var bounds = new L.LatLngBounds([0, 0], [worldWidth,
    worldHeight]);
map.setMaxBounds(bounds);
map.fitBounds(bounds);

map.options.minZoom = map.getZoom();

var imageUrl = '/map'
var imageBounds = [[0, 0],
[worldWidth, worldHeight]];

var overlayMaps = {
    "Sites": sitesLayer,
    "World Constructions": constructionsLayer,
    "Mountain Peaks": mountainsLayer,
    "Landmasses": landmassesLayer,
    "Regions": regionsLayer,
    "Evilness": evilnessLayer,
};

L.control.layers(null, overlayMaps).addTo(map);

var imageLayer = L.imageOverlay(imageUrl, imageBounds, { opacity: 0.5 });
imageLayer.addTo(map);

// var opacitySlider = new L.Control.opacitySlider();
// map.addControl(opacitySlider);
// opacitySlider.setOpacityLayer(imageLayer);


var minx = 1000, miny = 1000, maxx = 0, maxy = 0;

function zoomTo(y, x, zoom) {
    x = worldWidth - x - 1;
    map.setView([x, y], zoom);
}

function zoom() {
    var sw = L.latLng(minx, miny), ne = L.latLng(maxx + 1, maxy + 1);
    var bounds = new L.LatLngBounds(sw, ne);
    console.log(sw, ne, bounds);
    map.fitBounds(bounds);
}


var siteOffset = 0.1;
var structureOffset = 0.35;
var battleOffset = 0.2;
var mountainOffset = -0.2;

function coordO(y, x, yo, xo) {
    var c = coord(y, x);
    return [c[0] + yo, c[1] + xo]
}

function square(y, x, o) {
    return [coordO(y, x, o, o), coordO(y, x, 1 - o, o), coordO(y, x, 1 - o, 1 - o), coordO(y, x, o, 1 - o)];
}

function attachTooltip(layer, tip) {
    layer.bindTooltip(tip, { direction: 'top' }).bindPopup(tip);
}

function addSite(name, y1, x1, y2, x2, color, glyph) {
    /* resize tiny sites like lairs */
    var MIN_SIZE = .3;
    if (y2 - y1 < MIN_SIZE) {
        y1 = (y1 + y2) / 2 - MIN_SIZE / 2;
        y2 = y1 + MIN_SIZE;
    }
    if (x2 - x1 < MIN_SIZE) {
        x1 = (x1 + x2) / 2 - MIN_SIZE / 2;
        x2 = x1 + MIN_SIZE;
    }
    /* TODO: use glyph of the site instead of a polygon? */
    var polygon = L.polygon(
        [coord(y1, x1), coord(y2, x1), coord(y2, x2), coord(y1, x2)], {
        color: color,
        opacity: 1, fillOpacity: 0.7,
        weight: 3
    }).addTo(sitesLayer);

    attachTooltip(polygon, name);
}

function addWc(name, y, x, color) {
    var polygon = L.polygon(square(y, x, structureOffset), {
        color: color,
        opacity: 1, fillOpacity: 0.7,
        weight: 3
    }).addTo(constructionsLayer);

    attachTooltip(polygon, name);
}

function addRegion(name, y1, x1, y2, x2, color) {
    x1--; y2++;
    var polygon = L.polygon(
        [coord(y1, x1), coord(y2, x1), coord(y2, x2), coord(y1, x2)], {
        color: color,
        opacity: 0.5, fillOpacity: 0.3,
        weight: 1
    }).addTo(landmassesLayer);

    attachTooltip(polygon, name);
}

function addMountain(name, y, x, color) {
    x = worldWidth - x - 1;
    var polygon = L.polygon(
        [[x + mountainOffset / 2, y + mountainOffset], [x + mountainOffset / 2, y + 1 - mountainOffset], [x + 1 - mountainOffset, y + 0.5]], {
        color: color,
        opacity: 1, fillOpacity: 0.7,
        weight: 3
    }).addTo(mountainsLayer);

    attachTooltip(polygon, name);

    minx = Math.min(x, minx);
    miny = Math.min(y, miny);
    maxx = Math.max(x, maxx);
    maxy = Math.max(y, maxy);
}

function coord(y, x) {
    x = worldWidth - x - 1;

    minx = Math.min(x, minx);
    miny = Math.min(y, miny);
    maxx = Math.max(x, maxx);
    maxy = Math.max(y, maxy);

    return [x, y];
}

function addBattle(name, y, x) {
    x = worldWidth - x - 1;
    var polygon = L.polygon(
        [[x + 0.5, y + battleOffset],
        [x + battleOffset, y + 0.5],
        [x + 0.5, y + 1 - battleOffset],
        [x + 1 - battleOffset, y + 0.5]], {
        color: '#f00',
        opacity: 1, fillOpacity: 0.7,
        weight: 3
    }).addTo(map);

    attachTooltip(polygon, name);

    minx = Math.min(x, minx);
    miny = Math.min(y, miny);
    maxx = Math.max(x, maxx);
    maxy = Math.max(y, maxy);
}