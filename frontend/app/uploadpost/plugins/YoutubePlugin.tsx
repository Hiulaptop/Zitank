import { useState } from "react";
import { useLexicalComposerContext } from "@lexical/react/LexicalComposerContext";
import { $insertNodes } from "lexical";
import { $createYoutubeNode } from "../nodes/YoutubeNode";
import { Youtube } from "lucide-react";

export default function YoutubePlugin() {
  const [isOpen, setIsOpen] = useState(false);
  const [url, setURL] = useState("");
  const [editor] = useLexicalComposerContext();

  const onEmbed = () => {
    if (!url) return;
    const match =
      /^.*(youtu\.be\/|v\/|u\/\w\/|embed\/|watch\?v=|&v=)([^#&?]*).*/.exec(url);

    const id = match && match?.[2]?.length === 11 ? match?.[2] : null;
    if (!id) return;

    editor.update(() => {
      const node = $createYoutubeNode({ id });
      $insertNodes([node]);
    });

    setURL("");
    setIsOpen(false);
  };

  return (
    <div>
      {/* Button to open modal */}
      <button
        className="p-2 rounded-full hover:bg-red-200 text-red-600"
        onClick={() => setIsOpen(true)}
        aria-label="Embed Youtube Video"
      >
        <Youtube className="w-5 h-5" />
      </button>

      {/* Modal */}
      {isOpen && (
            <div className="fixed z-10 w-96 h-96">
              <input
                type="text"
                value={url}
                onChange={(e) => setURL(e.target.value)}
                placeholder="Paste Youtube URL"
                className="w-full mt-3 p-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-blue-500 z-50"
                autoFocus={true}
              />

              {/* Modal Footer */}
              <div className="flex justify-end mt-4 space-x-2">
                <button
                  onClick={() => setIsOpen(false)}
                  className="px-4 py-2 text-gray-600 bg-gray-100 rounded-md hover:bg-gray-200"
                >
                  Cancel
                </button>
                <button
                  onClick={onEmbed}
                  disabled={!url}
                  className={`px-4 py-2 text-white rounded-md ${
                    url
                      ? "bg-blue-600 hover:bg-blue-700"
                      : "bg-gray-400 cursor-not-allowed"
                  }`}
                >
                  Embed
                </button>
              </div>
            </div>
      )}

    </div>
  );
}
