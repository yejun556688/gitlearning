/*
 * Copyright 2013-2016 Canonical Ltd.
 *
 * This file is part of webbrowser-app.
 *
 * webbrowser-app is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; version 3.
 *
 * webbrowser-app is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

import QtQuick 2.4
import QtQuick.Window 2.2
import com.canonical.Oxide 1.12 as Oxide
import Ubuntu.Components 1.3
import Ubuntu.Components.Popups 1.3
import "." // QTBUG-34418

Oxide.WebView {
    id: _webview

    /**
     * Client overridable function called before the default treatment of a
     *  valid navigation request. This function can stop the navigation request
     *  if it sets the 'action' field of the request to IgnoreRequest.
     *
     */
    function navigationRequestedDelegate(request) { }

    context: SharedWebContext.sharedContext

    messageHandlers: [
        Oxide.ScriptMessageHandler {
            msgId: "scroll"
            contexts: ["oxide://selection/"]
            callback: function(msg, frame) {
                internal.dismissCurrentContextualMenu()
            }
        }
    ]

    onNavigationRequested: {
        request.action = Oxide.NavigationRequest.ActionAccept;
        navigationRequestedDelegate(request);
    }

    preferences.passwordEchoEnabled: Qt.inputMethod.visible

    popupMenu: ItemSelector02 {
        automaticOrientation: false
    }

    Item {
        id: contextualRectangle
        visible: false
        readonly property real locationBarOffset: _webview.locationBarController.height + _webview.locationBarController.offset
        x: internal.contextModel ? internal.contextModel.position.x : 0
        y: internal.contextModel ? internal.contextModel.position.y + locationBarOffset : 0
    }

    // XXX: This property is deprecated in favour of contextModel.
    property QtObject contextualData: QtObject {
        property url href
        property string title
        property url img

        function clear() {
            href = ''
            title = ''
            img = ''
        }
    }

    property var contextualActions // type: ActionList
    contextMenu: ActionSelectionPopover {
        objectName: "contextMenu"
        actions: contextualActions
        caller: contextualRectangle

        // Override default implementation to prevent context menu from stealing
        // active focus when shown (https://launchpad.net/bugs/1526884).
        function show() {
            visible = true
            __foreground.show()
        }

        Component.onCompleted: {
            internal.dismissCurrentContextualMenu()
            internal.contextModel = model
            var empty = true
            if (actions) {
                for (var i in actions.actions) {
                    if (actions.actions[i].enabled) {
                        empty = false
                        break
                    }
                }
            }
            if (empty) {
                internal.dismissCurrentContextualMenu()
            } else {
                contextualData.clear()
                contextualData.href = model.linkUrl
                contextualData.title = model.linkText
                if ((model.mediaType == Oxide.WebView.MediaTypeImage) && model.hasImageContents) {
                    contextualData.img = model.srcUrl
                }
                show()
            }
        }
        onVisibleChanged: {
            if (!visible) {
                internal.dismissCurrentContextualMenu()
            }
        }

        Binding {
            // Ensure the context menu doesn’t steal focus from
            // the webview when one of its actions is activated
            // (https://launchpad.net/bugs/1526884).
            target: __foreground
            property: "activeFocusOnPress"
            value: false
        }
    }
    readonly property QtObject contextModel: internal.contextModel

    property var selectionActions // type: ActionList
    onSelectionActionsChanged: console.warn("WARNING: the 'selectionActions' property is deprecated and ignored.")
    function copy() {
        console.warn("WARNING: the copy() function is deprecated and does nothing.")
    }

    touchSelectionController.handle: Image {
        objectName: "touchSelectionHandle"
        readonly property int handleOrientation: orientation
        width: units.gu(1.5)
        height: units.gu(1.5)
        source: "handle.png"
        Component.onCompleted: horizontalPaddingRatio = 0.5
    }

    UbuntuShape {
        objectName: "touchSelectionActions"
        // FIXME: hide contextual actions while resizing the
        // selection (needs an additional API in oxide?)
        visible: _webview.activeFocus && _webview.touchSelectionController.active && !selectionOutOfSight
        aspect: UbuntuShape.DropShadow
        backgroundColor: "white"
        readonly property int padding: units.gu(1)
        width: touchSelectionActionsRow.width + padding * 2
        height: childrenRect.height + padding * 2

        readonly property rect bounds: _webview.touchSelectionController.bounds
        readonly property bool selectionOutOfSight: (bounds.x > _webview.width) || ((bounds.x + bounds.width) < 0) || (bounds.y > _webview.height) || ((bounds.y + bounds.height) < 0)
        readonly property real handleHeight: units.gu(1.5)
        readonly property real spacing: units.gu(1)
        readonly property bool fitsBelow: (bounds.y + bounds.height + handleHeight + spacing + height) <= _webview.height
        readonly property bool fitsAbove: (bounds.y - spacing - height) >= (_webview.locationBarController.height + _webview.locationBarController.offset)
        readonly property real xCentered: bounds.x + (bounds.width - width) / 2
        x: ((xCentered >= 0) && ((xCentered + width) <= _webview.width))
            ? xCentered : (xCentered < 0) ? 0 : _webview.width - width
        y: fitsBelow ? (bounds.y + bounds.height + handleHeight + spacing)
                     : fitsAbove ? (bounds.y - spacing - height)
                                 : (_webview.height + _webview.locationBarController.height + _webview.locationBarController.offset - height) / 2

        ActionList {
            id: touchSelectionActions
            Action {
                name: "selectall"
                text: i18n.dtr('ubuntu-ui-toolkit', "Select All")
                iconName: "edit-select-all"
                enabled: _webview.editingCapabilities & Oxide.WebView.SelectAllCapability
                visible: enabled
                onTriggered: _webview.executeEditingCommand(Oxide.WebView.EditingCommandSelectAll)
            }
            Action {
                name: "cut"
                text: i18n.dtr('ubuntu-ui-toolkit', "Cut")
                iconName: "edit-cut"
                enabled: _webview.editingCapabilities & Oxide.WebView.CutCapability
                visible: enabled
                onTriggered: _webview.executeEditingCommand(Oxide.WebView.EditingCommandCut)
            }
            Action {
                name: "copy"
                text: i18n.dtr('ubuntu-ui-toolkit', "Copy")
                iconName: "edit-copy"
                enabled: _webview.editingCapabilities & Oxide.WebView.CopyCapability
                visible: enabled
                onTriggered: _webview.executeEditingCommand(Oxide.WebView.EditingCommandCopy)
            }
            Action {
                name: "paste"
                text: i18n.dtr('ubuntu-ui-toolkit', "Paste")
                iconName: "edit-paste"
                enabled: _webview.editingCapabilities & Oxide.WebView.PasteCapability
                visible: enabled
                onTriggered: _webview.executeEditingCommand(Oxide.WebView.EditingCommandPaste)
            }
        }

        Row {
            id: touchSelectionActionsRow
            x: parent.padding
            y: parent.padding
            width: {
                // work around what seems to be a bug in Row’s childrenRect.width
                var w = 0
                for (var i in visibleChildren) {
                    w += visibleChildren[i].width
                }
                return w
            }
            height: units.gu(6)

            Repeater {
                model: touchSelectionActions.actions.length
                AbstractButton {
                    objectName: "touchSelectionAction_" + action.name
                    anchors {
                        top: parent.top
                        bottom: parent.bottom
                    }
                    width: Math.max(units.gu(5), implicitWidth) + units.gu(2)
                    action: touchSelectionActions.actions[modelData]
                    styleName: "ToolbarButtonStyle"
                    activeFocusOnPress: false
                }
            }
        }
    }

    QtObject {
        id: internal
        property int lastLoadRequestStatus: -1
        property QtObject contextModel: null

        function dismissCurrentContextualMenu() {
            var model = contextModel
            contextModel = null
            if (model) {
                model.close()
            }
        }

        onContextModelChanged: if (!contextModel) _webview.contextualData.clear()
    }

    readonly property bool lastLoadSucceeded: internal.lastLoadRequestStatus === Oxide.LoadEvent.TypeSucceeded
    readonly property bool lastLoadStopped: internal.lastLoadRequestStatus === Oxide.LoadEvent.TypeStopped
    readonly property bool lastLoadFailed: internal.lastLoadRequestStatus === Oxide.LoadEvent.TypeFailed
    onLoadEvent: {
        if (!event.isError) {
            internal.lastLoadRequestStatus = event.type
        }
        internal.dismissCurrentContextualMenu()
    }

    readonly property int screenOrientation: Screen.orientation
    onScreenOrientationChanged: {
        internal.dismissCurrentContextualMenu()
    }

    onJavaScriptConsoleMessage: {
        if (_webview.incognito) {
            return
        }

        var msg = "[JS] (%1:%2) %3".arg(sourceId).arg(lineNumber).arg(message)
        if (level === Oxide.WebView.LogSeverityVerbose) {
            console.log(msg)
        } else if (level === Oxide.WebView.LogSeverityInfo) {
            console.info(msg)
        } else if (level === Oxide.WebView.LogSeverityWarning) {
            console.warn(msg)
        } else if ((level === Oxide.WebView.LogSeverityError) ||
                   (level === Oxide.WebView.LogSeverityErrorReport) ||
                   (level === Oxide.WebView.LogSeverityFatal)) {
            console.error(msg)
        }
    }
}
