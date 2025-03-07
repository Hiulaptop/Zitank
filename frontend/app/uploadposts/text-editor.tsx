'use client';

import type { JSX } from 'react';

import { AutoFocusPlugin } from '@lexical/react/LexicalAutoFocusPlugin';
import { CharacterLimitPlugin } from '@lexical/react/LexicalCharacterLimitPlugin';
import { CheckListPlugin } from '@lexical/react/LexicalCheckListPlugin';
import { ClearEditorPlugin } from '@lexical/react/LexicalClearEditorPlugin';
import { ClickableLinkPlugin } from '@lexical/react/LexicalClickableLinkPlugin';
import { CollaborationPlugin } from '@lexical/react/LexicalCollaborationPlugin';
import { useLexicalComposerContext } from '@lexical/react/LexicalComposerContext';
import { LexicalErrorBoundary } from '@lexical/react/LexicalErrorBoundary';
import { HashtagPlugin } from '@lexical/react/LexicalHashtagPlugin';
import { HistoryPlugin } from '@lexical/react/LexicalHistoryPlugin';
import { HorizontalRulePlugin } from '@lexical/react/LexicalHorizontalRulePlugin';
import { ListPlugin } from '@lexical/react/LexicalListPlugin';
import { PlainTextPlugin } from '@lexical/react/LexicalPlainTextPlugin';
import { RichTextPlugin } from '@lexical/react/LexicalRichTextPlugin';
import { SelectionAlwaysOnDisplay } from '@lexical/react/LexicalSelectionAlwaysOnDisplay';
import { TabIndentationPlugin } from '@lexical/react/LexicalTabIndentationPlugin';
import { TablePlugin } from '@lexical/react/LexicalTablePlugin';
import { useLexicalEditable } from '@lexical/react/useLexicalEditable';
import * as React from 'react';
import { useEffect, useState } from 'react';

import { useSettings } from './context/SettingsContext';
import { useSharedHistoryContext } from './context/SharedHistoryContext';
import ActionsPlugin from './plugins/ActionsPlugin';
import AutocompletePlugin from './plugins/AutocompletePlugin';
import AutoEmbedPlugin from './plugins/AutoEmbedPlugin';
import AutoLinkPlugin from './plugins/AutoLinkPlugin';
import CodeActionMenuPlugin from './plugins/CodeActionMenuPlugin';
import CodeHighlightPlugin from './plugins/CodeHighlightPlugin';
import CollapsiblePlugin from './plugins/CollapsiblePlugin';
import CommentPlugin from './plugins/CommentPlugin';
import ComponentPickerPlugin from './plugins/ComponentPickerPlugin';
import ContextMenuPlugin from './plugins/ContextMenuPlugin';
import DragDropPaste from './plugins/DragDropPastePlugin';
import DraggableBlockPlugin from './plugins/DraggableBlockPlugin';
import EmojiPickerPlugin from './plugins/EmojiPickerPlugin';
import EmojisPlugin from './plugins/EmojisPlugin';
import EquationsPlugin from './plugins/EquationsPlugin';
import FigmaPlugin from './plugins/FigmaPlugin';
import FloatingLinkEditorPlugin from './plugins/FloatingLinkEditorPlugin';
import FloatingTextFormatToolbarPlugin from './plugins/FloatingTextFormatToolbarPlugin';
import ImagesPlugin from './plugins/ImagesPlugin';
import InlineImagePlugin from './plugins/InlineImagePlugin';
import KeywordsPlugin from './plugins/KeywordsPlugin';
import { LayoutPlugin } from './plugins/LayoutPlugin/LayoutPlugin';
import LinkPlugin from './plugins/LinkPlugin';
import MarkdownShortcutPlugin from './plugins/MarkdownShortcutPlugin';
import { MaxLengthPlugin } from './plugins/MaxLengthPlugin';
import MentionsPlugin from './plugins/MentionsPlugin';
import PageBreakPlugin from './plugins/PageBreakPlugin';
import PollPlugin from './plugins/PollPlugin';
import ShortcutsPlugin from './plugins/ShortcutsPlugin';
import SpecialTextPlugin from './plugins/SpecialTextPlugin';
import SpeechToTextPlugin from './plugins/SpeechToTextPlugin';
import TabFocusPlugin from './plugins/TabFocusPlugin';
import TableCellActionMenuPlugin from './plugins/TableActionMenuPlugin';
import TableCellResizer from './plugins/TableCellResizer';
import TableHoverActionsPlugin from './plugins/TableHoverActionsPlugin';
import TableOfContentsPlugin from './plugins/TableOfContentsPlugin';
import ToolbarPlugin from './plugins/ToolbarPlugin';
import TreeViewPlugin from './plugins/TreeViewPlugin';
import TwitterPlugin from './plugins/TwitterPlugin';
import YouTubePlugin from './plugins/YouTubePlugin';
import ContentEditable from './ui/ContentEditable';

import { LexicalComposer } from '@lexical/react/LexicalComposer';

import { isDevPlayground } from './appSettings';
import { SharedHistoryContext } from './context/SharedHistoryContext';
import { ToolbarContext } from './context/ToolbarContext';
import logo from './images/logo.svg';
import PlaygroundNodes from './nodes/PlaygroundNodes';
import DocsPlugin from './plugins/DocsPlugin';
import PasteLogPlugin from './plugins/PasteLogPlugin';
import { TableContext } from './plugins/TablePlugin';
import TestRecorderPlugin from './plugins/TestRecorderPlugin';
import TypingPerfPlugin from './plugins/TypingPerfPlugin';
import Settings from './Settings';
import PlaygroundEditorTheme from './themes/PlaygroundEditorTheme';

const skipCollaborationInit =
    // @ts-expect-error
    window.parent != null && window.parent.frames.right === window;

export default function CoreEditor(): JSX.Element {
    const { historyState } = useSharedHistoryContext();
    const {
        settings: {
            isCollab,
            isAutocomplete,
            isMaxLength,
            isCharLimit,
            hasLinkAttributes,
            isCharLimitUtf8,
            isRichText,
            showTreeView,
            showTableOfContents,
            shouldUseLexicalContextMenu,
            shouldPreserveNewLinesInMarkdown,
            tableCellMerge,
            tableCellBackgroundColor,
            tableHorizontalScroll,
            shouldAllowHighlightingWithBrackets,
            selectionAlwaysOnDisplay,
        },
    } = useSettings();
    const isEditable = useLexicalEditable();
    const placeholder = isCollab
        ? 'Enter some collaborative rich text...'
        : isRichText
            ? 'Enter some rich text...'
            : 'Enter some plain text...';
    const [floatingAnchorElem, setFloatingAnchorElem] =
        useState<HTMLDivElement | null>(null);
    const [isSmallWidthViewport, setIsSmallWidthViewport] =
        useState<boolean>(false);
    const [editor] = useLexicalComposerContext();
    const [activeEditor, setActiveEditor] = useState(editor);
    const [isLinkEditMode, setIsLinkEditMode] = useState<boolean>(false);

    const onRef = (_floatingAnchorElem: HTMLDivElement) => {
        if (_floatingAnchorElem !== null) {
            setFloatingAnchorElem(_floatingAnchorElem);
        }
    };

    useEffect(() => {
        const updateViewPortWidth = () => {
            const isNextSmallWidthViewport =
                window.matchMedia('(max-width: 1025px)').matches;

            if (isNextSmallWidthViewport !== isSmallWidthViewport) {
                setIsSmallWidthViewport(isNextSmallWidthViewport);
            }
        };
        updateViewPortWidth();
        window.addEventListener('resize', updateViewPortWidth);

        return () => {
            window.removeEventListener('resize', updateViewPortWidth);
        };
    }, [isSmallWidthViewport]);

    return (
        <>
            {isRichText && (
                <ToolbarPlugin
                    editor={editor}
                    activeEditor={activeEditor}
                    setActiveEditor={setActiveEditor}
                    setIsLinkEditMode={setIsLinkEditMode}
                />
            )}
            {isRichText && (
                <ShortcutsPlugin
                    editor={activeEditor}
                    setIsLinkEditMode={setIsLinkEditMode}
                />
            )}
            <div
                className={`editor-container ${showTreeView ? 'tree-view' : ''} ${!isRichText ? 'plain-text' : ''
                    }`}>
                {isMaxLength && <MaxLengthPlugin maxLength={30} />}
                <DragDropPaste />
                <AutoFocusPlugin />
                {selectionAlwaysOnDisplay && <SelectionAlwaysOnDisplay />}
                <ClearEditorPlugin />
                <ComponentPickerPlugin />
                <EmojiPickerPlugin />
                <AutoEmbedPlugin />
                <MentionsPlugin />
                <EmojisPlugin />
                <HashtagPlugin />
                <KeywordsPlugin />
                <SpeechToTextPlugin />
                <AutoLinkPlugin />
                <CommentPlugin
                    providerFactory={undefined}
                />
                {isRichText ? (
                    <>
                        <HistoryPlugin externalHistoryState={historyState} />
                        <RichTextPlugin
                            contentEditable={
                                <div className="editor-scroller">
                                    <div className="editor" ref={onRef}>
                                        <ContentEditable placeholder={placeholder} />
                                    </div>
                                </div>
                            }
                            ErrorBoundary={LexicalErrorBoundary}
                        />
                        <MarkdownShortcutPlugin />
                        <CodeHighlightPlugin />
                        <ListPlugin />
                        <CheckListPlugin />
                        <TablePlugin
                            hasCellMerge={tableCellMerge}
                            hasCellBackgroundColor={tableCellBackgroundColor}
                            hasHorizontalScroll={tableHorizontalScroll}
                        />
                        <TableCellResizer />
                        <ImagesPlugin />
                        <InlineImagePlugin />
                        <LinkPlugin hasLinkAttributes={hasLinkAttributes} />
                        <PollPlugin />
                        <TwitterPlugin />
                        <YouTubePlugin />
                        <FigmaPlugin />
                        <ClickableLinkPlugin disabled={isEditable} />
                        <HorizontalRulePlugin />
                        <EquationsPlugin />
                        <TabFocusPlugin />
                        <TabIndentationPlugin maxIndent={7} />
                        <CollapsiblePlugin />
                        <PageBreakPlugin />
                        <LayoutPlugin />
                        {floatingAnchorElem && !isSmallWidthViewport && (
                            <>
                                <DraggableBlockPlugin anchorElem={floatingAnchorElem} />
                                <CodeActionMenuPlugin anchorElem={floatingAnchorElem} />
                                <FloatingLinkEditorPlugin
                                    anchorElem={floatingAnchorElem}
                                    isLinkEditMode={isLinkEditMode}
                                    setIsLinkEditMode={setIsLinkEditMode}
                                />
                                <TableCellActionMenuPlugin
                                    anchorElem={floatingAnchorElem}
                                    cellMerge={true}
                                />
                                <TableHoverActionsPlugin anchorElem={floatingAnchorElem} />
                                <FloatingTextFormatToolbarPlugin
                                    anchorElem={floatingAnchorElem}
                                    setIsLinkEditMode={setIsLinkEditMode}
                                />
                            </>
                        )}
                    </>
                ) : (
                    <>
                        <PlainTextPlugin
                            contentEditable={<ContentEditable placeholder={placeholder} />}
                            ErrorBoundary={LexicalErrorBoundary}
                        />
                        <HistoryPlugin externalHistoryState={historyState} />
                    </>
                )}
                {(isCharLimit || isCharLimitUtf8) && (
                    <CharacterLimitPlugin
                        charset={isCharLimit ? 'UTF-16' : 'UTF-8'}
                        maxLength={5}
                    />
                )}
                {isAutocomplete && <AutocompletePlugin />}
                <div>{showTableOfContents && <TableOfContentsPlugin />}</div>
                {shouldUseLexicalContextMenu && <ContextMenuPlugin />}
                {shouldAllowHighlightingWithBrackets && <SpecialTextPlugin />}
                <ActionsPlugin
                    isRichText={isRichText}
                    shouldPreserveNewLinesInMarkdown={shouldPreserveNewLinesInMarkdown}
                />
            </div>
            {showTreeView && <TreeViewPlugin />}
        </>
    );
}

export function Editor(): JSX.Element | null {
    const {
        settings: { emptyEditor },
    } = useSettings();

    const initialConfig = {
        editorState: undefined,
        namespace: 'Playground',
        nodes: [...PlaygroundNodes],
        onError: (error: Error) => {
            throw error;
        },
        theme: PlaygroundEditorTheme,
    };
    return (
        <LexicalComposer initialConfig={initialConfig}>
            <SharedHistoryContext>
                <TableContext>
                    <ToolbarContext>
                        <div className="editor-shell">
                            <CoreEditor />
                        </div>
                        <Settings />
                    </ToolbarContext>
                </TableContext>
            </SharedHistoryContext>
        </LexicalComposer>
    );
}
