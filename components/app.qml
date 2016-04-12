import QtQuick 2.5
import QtQuick.Controls 1.4
import QtQuick.Layouts 1.0
import QtWebEngine 1.0

ApplicationWindow {
    visible: true
    title: "Moa"
    width: 800
    height: 600

    MenuBar {
        Menu {
            title: "File"
            MenuItem {
                text: "Open..."
            }
            MenuItem {
                text: "Save"
            }
        }

        Menu {
            title: "Help"
            MenuItem {
                text: "Markdown Syntax"
            }
            MenuItem {
                objectName: "menu:help:about"
                text: "About"
            }
        }
    }

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
