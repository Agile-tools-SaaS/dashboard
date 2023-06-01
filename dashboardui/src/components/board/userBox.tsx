export const UsersBox = (users: { name: string; color: string }[]) => {
  return (
    <div
      style={{
        position: "absolute",
        width: "fit-content",
        bottom: 20,
        right: 20,
        boxShadow: "0px 0px 2px rgba(40,40,40,0.2)",
        padding: 10,
        background: "white",
        borderRadius: 10,
        display: "flex",
        flexDirection: "column",
        gap: 5,
      }}
    >
      {Object.values(users).map((user: any) => (
        <div
          style={{
            padding: 2.5,
            paddingLeft: 10,
            paddingRight: 10,
            borderRadius: 10,
            background: user.color,
          }}
        >
          <span style={{ color: "white" }}>{user.name}</span>
        </div>
      ))}
    </div>
  );
};
