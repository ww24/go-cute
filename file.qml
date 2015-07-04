import QtQuick 2.0

Rectangle {
  width: line.width
  height: line.height * 4

  Image {
    id: line
    source: "image://imgprv/0"
  }

  Image {
    id: image1
    anchors.top: line.bottom
    source: "image://imgprv/1"
  }

  Image {
    id: image2
    anchors.top: image1.bottom
    source: "image://imgprv/2"
  }

  Image {
    id: image3
    anchors.top: image2.bottom
    source: "image://imgprv/3"
  }
}
