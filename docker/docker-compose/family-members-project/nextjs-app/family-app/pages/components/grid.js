import { Box } from "@mui/system";
import { Button, Grid } from "@mui/material";

export default function MyGrid({ handleClick }) {
  return (
    <center>
      <Box sx={{ width: "100%" }}>
        <Grid
          container
          justifyContent="center"
          alignItems="center"
          rowSpacing={1}
          columnSpacing={{ xs: 1, sm: 2, md: 3 }}
        >
          <Grid p xs={6}>
            <Button onClick={() => handleClick()}>
              Select All Family Members
            </Button>
          </Grid>
          <Grid p xs={6}>
            <Button>Select Oldest Family Member</Button>{" "}
          </Grid>
          <Grid p xs={6}>
            <Button>Select Youngest Family Member</Button>{" "}
          </Grid>
          <Grid p xs={6}>
            <Button>Select Middle Family Member</Button>
          </Grid>
        </Grid>
      </Box>
    </center>
  );
}
