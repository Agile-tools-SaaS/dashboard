import styles from "@/styles/board/container.module.scss";

export const BoardContainer = ({ children }: { children: JSX.Element }) => {
  return <div className={styles["board-container"]}>{children}</div>;
};
