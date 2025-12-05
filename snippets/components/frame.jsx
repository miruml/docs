export const Frame = ({
    src,
    alt = "Framed content",
    background,
    radius = "10px",
    padding = "20px",
    imgRadius = "6px",
}) => {
    const image = (
        <img
            src={src}
            alt={alt}
            style={{
                display: "flex",
                width: "100%",
                borderRadius: imgRadius,
                margin: 0,

            }}
        />
    );

    return (

        <div
            style={{
                backgroundImage: background ? `url(${background})` : undefined,
                backgroundSize: "cover",

                borderRadius: radius,
                border: 0,
                padding: padding,
                overflow: "hidden",
            }}
        >
            {image}
        </div>
    );
};
