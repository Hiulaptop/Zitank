'use client';

import { LexicalComposer } from "@lexical/react/LexicalComposer";
import { RichTextPlugin } from "@lexical/react/LexicalRichTextPlugin";
import { ContentEditable } from "@lexical/react/LexicalContentEditable";
import { useEffect } from "react";
import { useLexicalComposerContext } from "@lexical/react/LexicalComposerContext";
import { LexicalErrorBoundary } from "@lexical/react/LexicalErrorBoundary";

const LoadContentFromProps = ({ data }: { data: string }) => {
  const [editor] = useLexicalComposerContext();

  useEffect(() => {
    const editorState = editor.parseEditorState(data);
    editor.setEditorState(editorState);
  }, [editor, data]);

  return null;
};

const editorConfig = {
  namespace: 'need to edit',
  theme: {}, 
  editable: false, 
  onError: (error: any) => console.error(error),
  nodes: [],
};

export function ReadonlyEditor({ data }: { data: string }) {
  return (
    <LexicalComposer initialConfig={editorConfig}>
      <RichTextPlugin
              contentEditable={<ContentEditable className="border p-4 bg-gray-100 cursor-default" />}
              placeholder={<div className="text-gray-400">Không có nội dung</div>}
              ErrorBoundary={LexicalErrorBoundary}
        />
      <LoadContentFromProps data={data} />
    </LexicalComposer>
  );
}
