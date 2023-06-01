import { GiArrowCursor } from "react-icons/gi";
import styles from "@/styles/mouse/pointer.module.scss";
import { motion } from "framer-motion";
import { useEffect } from "react";
export const Pointer = ({
  coordinates,
  color,
  user,
  otherUser,
}: {
  coordinates: { x: number; y: number };
  color: string;
  user: string;
  otherUser?: boolean;
}) => {
  return (
    <motion.div
      className={styles.pointer}
      data-other-user={otherUser}
      style={{
        left: coordinates.x as number,
        top: coordinates.y as number,
        transitionDuration: otherUser ? "0.5s" : "0",
        transitionTimingFunction: "ease",
      }}
    >
      <GiArrowCursor size={18} color={color} stroke="white" strokeWidth={20} />

      {otherUser && (
        <span style={{ background: color }} className={styles["user-tag"]}>
          {user}
        </span>
      )}
    </motion.div>
  );
};
