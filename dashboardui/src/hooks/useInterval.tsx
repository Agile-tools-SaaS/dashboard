import { useEffect, useState } from "react";

export const useInterval = (cb: () => void, ms: number) => {
  const [time, setTime] = useState(0);

  useEffect(() => {
    const timer = setTimeout(() => {
      setTime(time + 1);
      cb();
    }, ms);
    return () => {
      clearTimeout(timer);
    };
  }, [time]);
};
