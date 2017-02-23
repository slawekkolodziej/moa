package app

const componentsAppQml = `
import QtQuick 2.5
import QtQuick.Controls 1.4
import QtQuick.Dialogs 1.0
import QtQuick.Layouts 1.0
import QtWebEngine 1.0

ApplicationWindow {
    visible: true
    title: "Moa"
    width: 800
    height: 600

    FileDialog {
        id: openFile
        objectName: "openFile"
        title: "Choose a file"
    }

    FileDialog {
        id: saveFile
        objectName: "saveFile"
        title: "Choose a file"
        selectExisting: false
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
`

const componentsAboutQml = `
import QtQuick 2.5
import QtQuick.Controls 1.4

ApplicationWindow {
    visible: true
    width: 200
    height: 200

    Text {
        text: "Some really useful stuff should go here... (five)"
    }
}
`

const componentsMenubarQml = `
import QtQuick 2.5
import QtQuick.Controls 1.4

MenuBar {
    Menu {
        title: "File"
        MenuItem {
            objectName: "menu:file:open"
            text: "Open..."
        }
        MenuItem {
            objectName: "menu:file:save"
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


// FileDialog {
//     id: fileDialog
//     objectName: "fileDialog"
//     title: "Please choose a file"
//     folder: shortcuts.home
//     onAccepted: {
//         console.log("You chose: " + fileDialog.fileUrls)
//     }
//     onRejected: {
//         console.log("Canceled")
//     }
// }
`