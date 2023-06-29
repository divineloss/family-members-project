import * as React from "react";
import { useState } from "react";
import Container from "@mui/material/Container";
import { Box } from "@mui/system";
import { Grid, Paper, Typography } from "@mui/material";
import Divider from "@mui/material/Divider";
import MyGrid from "./components/grid";
import { getAllMembers } from "./api/apiCalls";

export default function Home() {
  const [results, setResults] = useState([
    {
      id: 1,
      name: "empty",
    },
  ]);

  const handleAllMemClick = async () => {
    const res = await getAllMembers();
    if (res) {
      console.log(res);
      setResults(res);
    }
  };
  return (
    <Container maxWidth="sm">
      <Box justifyContent="center" sx={{ my: 4 }}>
        <Typography align="center" variant="h4" component="h1" gutterBottom>
          Family-App
        </Typography>
        <Divider />
        <MyGrid handleClick={handleAllMemClick} />
        <Divider />
        <Typography align="center" variant="h4" component="h1" gutterBottom>
          Results:
        </Typography>
        {results.map((k, v) => (
          <p className="text-xl text-center" key={k.id}>
            {k.name}
          </p>
        ))}
      </Box>
    </Container>
  );
}
