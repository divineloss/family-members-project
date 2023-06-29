import * as React from "react";
import Container from "@mui/material/Container";
import { Box } from "@mui/system";
import { Typography } from "@mui/material";

export default function Home() {
  return (
    <Container maxWidth="sm">
      <Box justifyContent="center" sx={{ my: 4 }}>
        <Typography align="center" variant="h4" component="h1" gutterBottom>
          Family-App
        </Typography>
      </Box>
    </Container>
  );
}
