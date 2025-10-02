import React from "react";
import { cn } from "@/lib/utils";

export default function Animation() {
  return (
    <div className="fixed inset-0 bg-black w-full h-full z-0">
      <div
        className={cn(
          "pointer-events-none absolute inset-0 [background-size:40px_40px] select-none",
          "[background-image:linear-gradient(to_right,#171717_1px,transparent_1px),linear-gradient(to_bottom,#171717_1px,transparent_1px)]"
        )}
      />
    </div>
  );
}
