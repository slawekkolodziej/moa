import QtQuick 2.5
import QtQuick.Controls 1.1
import QtQuick.Layouts 1.0
import QtWebEngine 1.0

ApplicationWindow {
    visible: true
    title: "Moa"
    width: 800
    height: 600

    SplitView {
        anchors.fill: parent

        TextEdit {
            id: source
            objectName: "source"
            width: parent.width * 0.5
        }

        Item {
            WebEngineView {
                id: target
                objectName: "target"
                width: source.width
                anchors.fill: parent
            }
        }
    }
}
