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
        Component.onCompleted: console.log(file.content)
    }

    SplitView {
        anchors.fill: parent

        TextArea {
            id: source
            width: parent.width * 0.5
            text: file.content
        }

        WebView {
            id: output
            width: source.width
            // html: "<p>This is <b>HTML</b>."
            loadHtml: "http://www.google.com"
        }
    }
}