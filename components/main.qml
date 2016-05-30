import QtQuick 2.5
import QtQuick.Controls 1.4
import QtQuick.Dialogs 1.0
import QtQuick.Layouts 1.0
import QtWebEngine 1.0

ApplicationWindow {
    title: "Moa"
    width: 100
    height: 100

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

    FileDialog {
        id: fileDialog
        objectName: "fileDialog"
        title: "Please choose a file"
        folder: shortcuts.home
        onAccepted: {
            console.log("You chose: " + fileDialog.fileUrls)
        }
        onRejected: {
            console.log("Canceled")
        }
    }
}
