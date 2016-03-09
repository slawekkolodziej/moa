import QtQuick 2.2
import QtQuick.Controls 1.1
import QtQuick.Layouts 1.0
import QtWebKit 3.0

ApplicationWindow {
    visible: true
    title: "Moa"
    width: 600
    height: 400

    Item {
        property alias text: source.text
    }

    SplitView {
        anchors.fill: parent

        TextEdit {
            id: source
            objectName: "source"
            width: parent.width * 0.5
        }

        WebView {
            id: target
            objectName: "target"
            width: source.width
        }
    }
}
