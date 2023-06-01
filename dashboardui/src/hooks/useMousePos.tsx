import { useEffect, useState } from "react";

interface mouseCoordinates {
  x: number;
  y: number;
}

export function useMousePos() {
  const [mousePos, setMousePos] = useState<mouseCoordinates>({ x: 0, y: 0 });

  useEffect(() => {
    const handleMouseMove = (e: MouseEvent) => {
      setMousePos({ x: e.clientX, y: e.clientY });
    };

    window.addEventListener("mousemove", handleMouseMove);

    return () => {
      window.removeEventListener("mousemove", handleMouseMove);
    };
  }, []);

  return mousePos;
}
